package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID   string
	Name string
}

type Split struct {
	OwedBy User
	OwedTo User
	Amount float64
}

type SplitStrategy interface {
	Split(amount float64, paidBy User, participants []User, meta map[string]any) ([]Split, error)
}

type EqualSplitStrategy struct{}

func (EqualSplitStrategy) Split(amount float64, paidBy User, participants []User, _ map[string]any) ([]Split, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	n := float64(len(participants))
	if n == 0 {
		return nil, errors.New("no participants")
	}
	share := round2(amount / n)

	residual := round2(amount - share*n)
	var splits []Split
	for _, u := range participants {
		amt := share
		if residual != 0 && u.ID == paidBy.ID {
			amt += residual
		}
		if u.ID == paidBy.ID {
			continue
		}
		splits = append(splits, Split{OwedBy: u, OwedTo: paidBy, Amount: amt})
	}
	return splits, nil
}

type ExactSplitStrategy struct{}

func (ExactSplitStrategy) Split(amount float64, paidBy User, participants []User, meta map[string]any) ([]Split, error) {
	sharesRaw, ok := meta["shares"]
	if !ok {
		return nil, errors.New("missing shares map in meta")
	}
	shares, ok := sharesRaw.(map[string]float64)
	if !ok {
		return nil, errors.New("shares must be map[string]float64")
	}
	sum := 0.0
	for _, v := range shares {
		sum += v
	}
	if math.Abs(sum-amount) > 0.01 {
		return nil, fmt.Errorf("exact shares (%.2f) do not sum to amount %.2f", sum, amount)
	}
	var splits []Split
	for _, u := range participants {
		share := round2(shares[u.ID])
		if u.ID == paidBy.ID || share == 0 {
			continue
		}
		splits = append(splits, Split{OwedBy: u, OwedTo: paidBy, Amount: share})
	}
	return splits, nil
}

type PercentSplitStrategy struct{}

func (PercentSplitStrategy) Split(amount float64, paidBy User, participants []User, meta map[string]any) ([]Split, error) {
	percRaw, ok := meta["percents"]
	if !ok {
		return nil, errors.New("missing percents map in meta")
	}
	percents, ok := percRaw.(map[string]float64)
	if !ok {
		return nil, errors.New("percents must be map[string]float64")
	}
	sum := 0.0
	for _, v := range percents {
		sum += v
	}
	if math.Abs(sum-100.0) > 0.01 {
		return nil, fmt.Errorf("percentages (%.2f) must sum to 100", sum)
	}
	var splits []Split
	for _, u := range participants {
		pct := percents[u.ID]
		share := round2(amount * pct / 100.0)
		if u.ID == paidBy.ID || share == 0 {
			continue
		}
		splits = append(splits, Split{OwedBy: u, OwedTo: paidBy, Amount: share})
	}
	return splits, nil
}

type StrategyType string

const (
	StrategyEqual   StrategyType = "EQUAL"
	StrategyExact   StrategyType = "EXACT"
	StrategyPercent StrategyType = "PERCENT"
)

func SplitStrategyFactory(t StrategyType) (SplitStrategy, error) {
	switch t {
	case StrategyEqual:
		return EqualSplitStrategy{}, nil
	case StrategyExact:
		return ExactSplitStrategy{}, nil
	case StrategyPercent:
		return PercentSplitStrategy{}, nil
	default:
		return nil, fmt.Errorf("unknown strategy %s", t)
	}
}

type ExpenseListener interface {
	OnExpenseRecorded(expense Expense, splits []Split)
}

// ExpensePublisher holds subscribers and broadcasts events.
type ExpensePublisher struct {
	mu          sync.RWMutex
	subscribers []ExpenseListener
}

func (p *ExpensePublisher) Register(l ExpenseListener) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscribers = append(p.subscribers, l)
}

func (p *ExpensePublisher) Publish(exp Expense, splits []Split) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, l := range p.subscribers {
		l.OnExpenseRecorded(exp, splits)
	}
}

type Ledger struct {
	mu       sync.RWMutex
	balances map[string]map[string]float64
}

func NewLedger() *Ledger {
	return &Ledger{balances: make(map[string]map[string]float64)}
}

func (l *Ledger) OnExpenseRecorded(_ Expense, splits []Split) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, s := range splits {
		l.addDebt(s.OwedBy.ID, s.OwedTo.ID, s.Amount)
	}
}

func (l *Ledger) addDebt(from, to string, amt float64) {
	if amt == 0 || from == to {
		return
	}

	if _, ok := l.balances[from]; !ok {
		l.balances[from] = make(map[string]float64)
	}
	if _, ok := l.balances[to]; !ok {
		l.balances[to] = make(map[string]float64)
	}

	if rev, ok := l.balances[to][from]; ok {
		if rev > amt {
			l.balances[to][from] = round2(rev - amt)
			return
		}
		amt = round2(amt - rev)
		delete(l.balances[to], from)
	}
	l.balances[from][to] = round2(l.balances[from][to] + amt)
	if l.balances[from][to] == 0 {
		delete(l.balances[from], to)
	}
}

func (l *Ledger) Pay(payerID, payeeID string, amount float64) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	curr := l.balances[payerID][payeeID]
	if amount <= 0 || amount > curr+0.001 {
		return fmt.Errorf("invalid payment amount")
	}
	l.balances[payerID][payeeID] = round2(curr - amount)
	if l.balances[payerID][payeeID] == 0 {
		delete(l.balances[payerID], payeeID)
	}
	return nil
}

func (l *Ledger) String() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	var sb strings.Builder
	for debtor, m := range l.balances {
		for creditor, amt := range m {
			sb.WriteString(fmt.Sprintf("%s owes %s: ₹%.2f\n", debtor, creditor, amt))
		}
	}
	if sb.Len() == 0 {
		return "Ledger is clean 🎉"
	}
	return sb.String()
}

type Expense struct {
	ID           string
	Description  string
	Amount       float64
	PaidBy       User
	Participants []User
	Strategy     SplitStrategy
	Meta         map[string]any
	RecordedAt   time.Time
}

type ExpenseService struct {
	publisher *ExpensePublisher
}

func NewExpenseService(pub *ExpensePublisher) *ExpenseService {
	return &ExpenseService{publisher: pub}
}

func (es *ExpenseService) AddExpense(desc string, amount float64, paidBy User, participants []User,
	strategyType StrategyType, meta map[string]any) error {

	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if len(participants) == 0 {
		return errors.New("participants cannot be empty")
	}
	// Ensure paidBy is in participants
	found := false
	for _, u := range participants {
		if u.ID == paidBy.ID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("payer must be part of participants")
	}

	strategy, err := SplitStrategyFactory(strategyType)
	if err != nil {
		return err
	}
	splits, err := strategy.Split(amount, paidBy, participants, meta)
	if err != nil {
		return err
	}

	exp := Expense{
		ID:           fmt.Sprintf("exp_%d", time.Now().UnixNano()),
		Description:  desc,
		Amount:       amount,
		PaidBy:       paidBy,
		Participants: participants,
		Strategy:     strategy,
		Meta:         meta,
		RecordedAt:   time.Now(),
	}

	es.publisher.Publish(exp, splits)
	return nil
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

func main() {
	alice := User{ID: "alice", Name: "Alice"}
	bob := User{ID: "bob", Name: "Bob"}
	charlie := User{ID: "charlie", Name: "Charlie"}

	// Prepare core pieces
	publisher := &ExpensePublisher{}
	ledger := NewLedger()
	publisher.Register(ledger)
	service := NewExpenseService(publisher)

	// 1) Equal split dinner
	_ = service.AddExpense(
		"Dinner",
		1200,
		alice,
		[]User{alice, bob, charlie},
		StrategyEqual,
		nil,
	)

	// 2) Bob paid cab: Bob & Charlie exact split
	_ = service.AddExpense(
		"Cab ride",
		500,
		bob,
		[]User{bob, charlie},
		StrategyExact,
		map[string]any{"shares": map[string]float64{
			"bob":     100, // Bob owes himself nothing, but we include for completeness
			"charlie": 400,
		}},
	)

	// 3) Charlie paid snacks: percentage split
	_ = service.AddExpense(
		"Snacks",
		300,
		charlie,
		[]User{alice, bob, charlie},
		StrategyPercent,
		map[string]any{"percents": map[string]float64{
			"alice":   40,
			"bob":     40,
			"charlie": 20,
		}},
	)

	fmt.Println("=== Ledger after expenses ===")
	fmt.Print(ledger)

	// Alice settles ₹400 she owes Bob
	if err := ledger.Pay("alice", "bob", 400); err != nil {
		fmt.Println("Payment error:", err)
	}

	fmt.Println("\n=== Ledger after settlement ===")
	fmt.Print(ledger)
}
