package model

import (
	"math"
	"sort"
	"sync"
)

type BreakingStrategy interface {
	Break(full Container) []Container
}

type ByColumn struct {
	BreakingStrategy
}

func (c *ByColumn) Break(full Container) []Container {
	half := math.Round(full.Length / 2)
	f := full
	f.Length = half

	s := full
	s.Length = half
	s.LeftBottomCorner = Point{X: full.LeftBottomCorner.X, Y: full.LeftBottomCorner.Y + half, Z: full.LeftBottomCorner.Z}

	return []Container{f, s}
}

type Pool struct {
	mu    sync.Mutex
	items []Item // this is a sorted list of item by volume, desc
}

func NewPool(i []Item) *Pool {
	sort.Slice(i, func(first, second int) bool {
		return i[first].Volume() > i[second].Volume()
	})
	p := &Pool{items: i}
	p.mu.Lock()
	return p
}

func (p *Pool) AllowTake() {
	p.mu.Unlock()
}

func (p *Pool) SafeTakeItem(volume float64) (bool, Item) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.items) == 0 {
		return false, Item{}
	}

	selected := p.items[0]
	if selected.Volume() <= volume {
		p.items = p.items[1:]
		return true, selected
	}

	return false, Item{}
}
