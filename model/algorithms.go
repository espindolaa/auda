package model

import "sync"

type BreakingStrategy interface {
	Break(full Container) []Container
}

type ByColumn struct {
	BreakingStrategy
}

func (c *ByColumn) Break(full Container) []Container {
	half := full.Length / 2
	f := full
	f.Length = half

	s := full
	s.Length = half
	s.LeftBottomPoint = Point{X: full.LeftBottomPoint.X, Y: full.LeftBottomPoint.Y + half, Z: full.LeftBottomPoint.Z}

	return []Container{f, s}
}

type Pool struct {
	mu    sync.Mutex
	items []Item // this is a sorted list of item by voulme, desc
}

func NewPool(items []Item) *Pool {
	// todo: sort items (merge sort)
	return &Pool{items: items}
}

func (p *Pool) SafeTakeItem(volume int) (bool, Item) {
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
