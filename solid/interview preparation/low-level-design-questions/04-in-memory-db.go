// kvstore.go
package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"
)

type entry struct {
	value  any
	expiry time.Time
}

type EvictionPolicy interface {
	OnPut(key string)
	OnGet(key string)
	OnDelete(key string)
	Evict(keys map[string]entry)
}

type lruPolicy struct {
	capacity int
	ll       *list.List
	nodes    map[string]*list.Element
}

func NewLRUPolicy(cap int) EvictionPolicy {
	return &lruPolicy{
		capacity: cap,
		ll:       list.New(),
		nodes:    make(map[string]*list.Element),
	}
}

func (p *lruPolicy) touch(key string) {
	if node, ok := p.nodes[key]; ok {
		p.ll.MoveToFront(node)
	}
}

func (p *lruPolicy) OnPut(key string) {
	if node, ok := p.nodes[key]; ok {
		p.ll.MoveToFront(node)
		return
	}
	p.nodes[key] = p.ll.PushFront(key)
}

func (p *lruPolicy) OnGet(key string) { p.touch(key) }

func (p *lruPolicy) OnDelete(key string) {
	if node, ok := p.nodes[key]; ok {
		p.ll.Remove(node)
		delete(p.nodes, key)
	}
}

func (p *lruPolicy) Evict(keys map[string]entry) {
	for len(keys) > p.capacity {
		lruElem := p.ll.Back()
		if lruElem == nil {
			return
		}
		lruKey := lruElem.Value.(string)
		delete(keys, lruKey)
		p.ll.Remove(lruElem)
		delete(p.nodes, lruKey)
	}
}

var (
	ErrKeyNotFound   = errors.New("key not found")
	ErrNegativeTTL   = errors.New("ttl must be non‑negative")
	ErrInvalidCap    = errors.New("capacity must be positive")
	ErrEvictionNil   = errors.New("eviction policy cannot be nil")
	ErrNilStoreValue = errors.New("value cannot be nil")
)

func validateTTL(ttl time.Duration) error {
	if ttl < 0 {
		return ErrNegativeTTL
	}
	return nil
}

type Store interface {
	Put(key string, val any, ttl time.Duration) error
	Get(key string) (any, error)
	Delete(key string) error
	Size() int
}

type inMemStore struct {
	mu      sync.RWMutex
	data    map[string]entry
	evictor EvictionPolicy
}

func NewInMemoryStore(capacity int, ev EvictionPolicy) (Store, error) {
	if capacity <= 0 {
		return nil, ErrInvalidCap
	}
	if ev == nil {
		return nil, ErrEvictionNil
	}
	return &inMemStore{
		data:    make(map[string]entry),
		evictor: ev,
	}, nil
}

func (s *inMemStore) Put(key string, val any, ttl time.Duration) error {
	if err := validateTTL(ttl); err != nil {
		return err
	}
	if val == nil {
		return ErrNilStoreValue
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	exp := time.Time{}
	if ttl > 0 {
		exp = time.Now().Add(ttl)
	}

	s.data[key] = entry{value: val, expiry: exp}
	s.evictor.OnPut(key)
	s.evictor.Evict(s.data)
	return nil
}

func (s *inMemStore) Get(key string) (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ent, ok := s.data[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	if ent.expiry.IsZero() || ent.expiry.After(time.Now()) {
		s.evictor.OnGet(key)
		return ent.value, nil
	}

	delete(s.data, key)
	s.evictor.OnDelete(key)
	return nil, ErrKeyNotFound
}

func (s *inMemStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		return ErrKeyNotFound
	}
	delete(s.data, key)
	s.evictor.OnDelete(key)
	return nil
}

func (s *inMemStore) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func main() {
	// Create store with capacity 2 & LRU eviction strategy
	store, err := NewInMemoryStore(2, NewLRUPolicy(2))
	if err != nil {
		panic(err)
	}

	fmt.Println(">>> Basic put / get")
	_ = store.Put("A", "🍎", 0)
	_ = store.Put("B", "🍌", 0)
	v, _ := store.Get("A")
	fmt.Println("A =", v) // 🍎

	fmt.Println(">>> Trigger eviction (capacity=2)")
	_ = store.Put("C", "🥥", 0) // LRU key "B" should be evicted
	_, err = store.Get("B")
	fmt.Println("B exists?", err == nil) // false

	fmt.Println(">>> TTL demo (2 s)")
	_ = store.Put("temp", "⏰", 2*time.Second)
	time.Sleep(3 * time.Second)
	_, err = store.Get("temp")
	fmt.Println("temp exists after 3s?", err == nil) // false

	fmt.Println(">>> Final size:", store.Size()) // should be 2 ("A" and "C")
}
