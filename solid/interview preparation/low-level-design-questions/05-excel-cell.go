package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var cellRefRe = regexp.MustCompile(`^[A-Z]+[0-9]+$`)

func cellName(row, col int) string {
	name := ""
	for col > 0 {
		col--
		name = string('A'+(col%26)) + name
		col /= 26
	}
	return fmt.Sprintf("%s%d", name, row)
}

type cell struct {
	raw   string
	val   any
	dirty bool

	deps       map[string]struct{}
	dependents map[string]struct{}
}

func newCell() *cell {
	return &cell{
		deps:       map[string]struct{}{},
		dependents: map[string]struct{}{},
	}
}

type Sheet struct {
	mu   sync.RWMutex
	grid map[string]*cell
}

func NewSheet() *Sheet {
	return &Sheet{grid: make(map[string]*cell)}
}

func (sh *Sheet) Set(ref, raw string) error {
	ref = strings.ToUpper(strings.TrimSpace(ref))
	if !cellRefRe.MatchString(ref) {
		return fmt.Errorf("invalid ref: %s", ref)
	}

	sh.mu.Lock()
	defer sh.mu.Unlock()

	c := sh.ensure(ref)

	for d := range c.deps {
		delete(sh.grid[d].dependents, ref)
	}
	c.deps = map[string]struct{}{}

	c.raw = strings.TrimSpace(raw)
	c.dirty = true

	if strings.HasPrefix(c.raw, "=") {
		for _, tok := range tokenScan(c.raw[1:]) {
			if cellRefRe.MatchString(tok) {
				c.deps[tok] = struct{}{}
				sh.ensure(tok).dependents[ref] = struct{}{}
			}
		}
	}

	sh.markDirty(ref)

	return nil
}

func (sh *Sheet) Get(ref string) (any, error) {
	sh.mu.RLock()
	defer sh.mu.RUnlock()

	ref = strings.ToUpper(strings.TrimSpace(ref))
	c, ok := sh.grid[ref]
	if !ok {
		return "", nil // blank cell
	}
	val, err := sh.eval(ref, map[string]struct{}{})
	return val, err
}

func (sh *Sheet) Print() {
	sh.mu.RLock()
	defer sh.mu.RUnlock()
	for r := 1; r <= 5; r++ {
		rowVals := make([]string, 5)
		for c := 1; c <= 5; c++ {
			ref := cellName(r, c)
			v, _ := sh.eval(ref, map[string]struct{}{})
			rowVals[c-1] = fmt.Sprintf("%v", v)
		}
		fmt.Println(strings.Join(rowVals, "\t"))
	}
}

func (sh *Sheet) ensure(ref string) *cell {
	if c, ok := sh.grid[ref]; ok {
		return c
	}
	c := newCell()
	sh.grid[ref] = c
	return c
}

func (sh *Sheet) markDirty(ref string) {
	c := sh.grid[ref]
	if !c.dirty {
		c.dirty = true
		for dep := range c.dependents {
			sh.markDirty(dep)
		}
	}
}

func (sh *Sheet) eval(ref string, path map[string]struct{}) (any, error) {
	c := sh.grid[ref]
	if !c.dirty {
		return c.val, nil
	}

	if _, ok := path[ref]; ok {
		c.val = "#CYCLE!"
		c.dirty = false
		return c.val, fmt.Errorf("cycle at %s", ref)
	}
	path[ref] = struct{}{}

	if !strings.HasPrefix(c.raw, "=") {
		if f, err := strconv.ParseFloat(c.raw, 64); err == nil {
			c.val = f
		} else {
			c.val = c.raw
		}
		c.dirty = false
		delete(path, ref)
		return c.val, nil
	}

	toks := tokenScan(c.raw[1:])
	if len(toks) != 3 {
		c.val = "#ERR!"
		c.dirty = false
		delete(path, ref)
		return c.val, fmt.Errorf("bad expr %s", c.raw)
	}
	left, op, right := toks[0], toks[1], toks[2]

	lv, err1 := sh.atom(left, path)
	rv, err2 := sh.atom(right, path)
	if err1 != nil || err2 != nil {
		c.val = "#ERR!"
	} else if lf, lok := lv.(float64); lok {
		if rf, rok := rv.(float64); rok {
			switch op {
			case "+":
				c.val = lf + rf
			case "-":
				c.val = lf - rf
			case "*":
				c.val = lf * rf
			case "/":
				if rf == 0 {
					c.val = "#DIV/0!"
				} else {
					c.val = lf / rf
				}
			default:
				c.val = "#ERR!"
			}
		} else {
			c.val = "#ERR!"
		}
	} else {
		c.val = "#ERR!"
	}
	c.dirty = false
	delete(path, ref)
	return c.val, nil
}

func (sh *Sheet) atom(tok string, path map[string]struct{}) (any, error) {
	if cellRefRe.MatchString(tok) {
		return sh.eval(tok, path)
	}
	if f, err := strconv.ParseFloat(tok, 64); err == nil {
		return f, nil
	}
	return nil, fmt.Errorf("invalid token %s", tok)
}

func tokenScan(expr string) []string {
	out := []string{}
	cur := ""
	for _, r := range expr {
		if r == '+' || r == '-' || r == '*' || r == '/' {
			if cur != "" {
				out = append(out, strings.TrimSpace(cur))
				cur = ""
			}
			out = append(out, string(r))
		} else {
			cur += string(r)
		}
	}
	if cur != "" {
		out = append(out, strings.TrimSpace(cur))
	}
	return out
}

func main() {
	s := NewSheet()
	_ = s.Set("A1", "5")
	_ = s.Set("B1", "10")
	_ = s.Set("C1", "=A1+B1")
	_ = s.Set("D1", "=C1*2")
	_ = s.Set("E1", "=D1/A1")

	fmt.Println("initial sheet")
	s.Print()

	// update A1; dependents auto-dirty
	_ = s.Set("A1", "20")
	fmt.Println("\nafter A1=20")
	s.Print()

	// introduce a cycle
	_ = s.Set("B2", "=C2")
	_ = s.Set("C2", "=B2")
	fmt.Println("\nwith cycle (B2<->C2)")
	s.Print()
}
