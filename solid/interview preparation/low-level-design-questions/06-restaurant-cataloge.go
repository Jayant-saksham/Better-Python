package main

import (
	"fmt"
	"sync"
)

/* ─────────────────────────── Domain structs ──────────────────────────── */

type Restaurant struct{ ID, Name string }
type Category struct{ ID, RestaurantID, Name string }
type Item struct {
	ID, CategoryID, Name string
	BasePrice            float64
}

type VariantGroup struct{ ID, ItemID, Name string }
type Variant struct {
	ID, GroupID, Name string
	Delta             float64
}

type AddOnGroup struct{ ID, ItemID, Name string }
type AddOn struct {
	ID, GroupID, Name string
	PricePerVariant   map[string]float64
}

type restRepo struct {
	m  map[string]Restaurant
	mu sync.RWMutex
}
type catRepo struct {
	m  map[string]Category
	mu sync.RWMutex
}
type itemRepo struct {
	m  map[string]Item
	mu sync.RWMutex
}
type vgRepo struct {
	m  map[string]VariantGroup
	mu sync.RWMutex
}
type varRepo struct {
	m  map[string]Variant
	mu sync.RWMutex
}
type agRepo struct {
	m  map[string]AddOnGroup
	mu sync.RWMutex
}
type aoRepo struct {
	m  map[string]AddOn
	mu sync.RWMutex
}

func newRestRepo() *restRepo { return &restRepo{m: map[string]Restaurant{}} }
func newCatRepo() *catRepo   { return &catRepo{m: map[string]Category{}} }
func newItemRepo() *itemRepo { return &itemRepo{m: map[string]Item{}} }
func newVgRepo() *vgRepo     { return &vgRepo{m: map[string]VariantGroup{}} }
func newVarRepo() *varRepo   { return &varRepo{m: map[string]Variant{}} }
func newAgRepo() *agRepo     { return &agRepo{m: map[string]AddOnGroup{}} }
func newAoRepo() *aoRepo     { return &aoRepo{m: map[string]AddOn{}} }

type CatalogService struct {
	rest *restRepo
	cat  *catRepo
	item *itemRepo
	vg   *vgRepo
	vars *varRepo
	ag   *agRepo
	ao   *aoRepo
}

func NewCatalogService() *CatalogService {
	return &CatalogService{
		rest: newRestRepo(), cat: newCatRepo(), item: newItemRepo(),
		vg: newVgRepo(), vars: newVarRepo(),
		ag: newAgRepo(), ao: newAoRepo(),
	}
}

func (s *CatalogService) CreateRestaurant(id, name string) {
	s.rest.mu.Lock()
	defer s.rest.mu.Unlock()
	s.rest.m[id] = Restaurant{id, name}
}

func (s *CatalogService) CreateCategory(id, restID, name string) {
	s.cat.mu.Lock()
	defer s.cat.mu.Unlock()
	s.cat.m[id] = Category{id, restID, name}
}

func (s *CatalogService) CreateItem(id, catID, name string, base float64) {
	s.item.mu.Lock()
	defer s.item.mu.Unlock()
	s.item.m[id] = Item{id, catID, name, base}
}

func (s *CatalogService) AddVariantGroup(id, itemID, name string) {
	s.vg.mu.Lock()
	defer s.vg.mu.Unlock()
	s.vg.m[id] = VariantGroup{id, itemID, name}
}

func (s *CatalogService) AddVariant(id, groupID, name string, delta float64) {
	s.vars.mu.Lock()
	defer s.vars.mu.Unlock()
	s.vars.m[id] = Variant{id, groupID, name, delta}
}

func (s *CatalogService) AddAddOnGroup(id, itemID, name string) {
	s.ag.mu.Lock()
	defer s.ag.mu.Unlock()
	s.ag.m[id] = AddOnGroup{id, itemID, name}
}

func (s *CatalogService) AddAddOn(id, groupID, name string, priceMap map[string]float64) {
	s.ao.mu.Lock()
	defer s.ao.mu.Unlock()
	s.ao.m[id] = AddOn{id, groupID, name, priceMap}
}

func (s *CatalogService) Price(itemID, variantID string, addOnIDs []string) (float64, error) {
	s.item.mu.RLock()
	it, ok := s.item.m[itemID]
	s.item.mu.RUnlock()
	if !ok {
		return 0, fmt.Errorf("item %s not found", itemID)
	}
	total := it.BasePrice

	if variantID != "" {
		s.vars.mu.RLock()
		v, vok := s.vars.m[variantID]
		s.vars.mu.RUnlock()
		if !vok {
			return 0, fmt.Errorf("variant %s missing", variantID)
		}
		total += v.Delta
	}

	for _, aoID := range addOnIDs {
		s.ao.mu.RLock()
		ao, aok := s.ao.m[aoID]
		s.ao.mu.RUnlock()
		if !aok {
			return 0, fmt.Errorf("addon %s missing", aoID)
		}
		if priceExtra, ok := ao.PricePerVariant[variantID]; ok {
			total += priceExtra
		}
	}

	return total, nil
}

func main() {
	s := NewCatalogService()

	s.CreateRestaurant("r1", "Pizza Hub")
	s.CreateCategory("cat1", "r1", "Pizzas")
	s.CreateItem("item1", "cat1", "Margherita", 200)

	s.AddVariantGroup("vgSize", "item1", "Size")
	s.AddVariant("vS", "vgSize", "Small", 0)
	s.AddVariant("vM", "vgSize", "Medium", 50)
	s.AddVariant("vL", "vgSize", "Large", 100)

	s.AddAddOnGroup("agCheese", "item1", "Cheese")
	s.AddAddOn("aoWhite", "agCheese", "White Cheese",
		map[string]float64{"vS": 30, "vM": 40, "vL": 50})
	s.AddAddOn("aoYellow", "agCheese", "Yellow Cheese",
		map[string]float64{"vS": 20, "vM": 30, "vL": 40})

	price0, _ := s.Price("item1", "vS", nil)
	fmt.Printf("Small, no add‑ons 👉 ₹%.0f\n", price0) // 200

	price1, _ := s.Price("item1", "vM", []string{"aoWhite"})
	fmt.Printf("Medium + White cheese 👉 ₹%.0f\n", price1) // 200+50+40

	price2, _ := s.Price("item1", "vL", []string{"aoWhite", "aoYellow"})
	fmt.Printf("Large + both cheese 👉 ₹%.0f\n", price2) // 200+100+50+40
}
