package auda

import (
	"fmt"
	"sync"
)

type Box struct {
	width  int
	length int
	height int
}

type Container struct {
	Box
	unsortedItems []Item
}

func (c *Container) pickItems(pool *Pool) {
	v := c.Volume()
	for {
		found, item := pool.SafeTakeItem(v)
		if !found {
			break
		}
		v = v - item.Volume()
		c.unsortedItems = append(c.unsortedItems, item)
	}
}

type Item struct {
	Box
}

func (b Box) Volume() int {
	return b.height * b.width * b.length
}

type Pool struct {
	mu    sync.Mutex
	items []Item // this is a sorted list of item by voulme, desc
}

func NewPool(items []Item) *Pool {
	// sort items (merge sort)
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

func main() {
	// todo: move to cmd input
	container := Container{Box{height: 10, width: 10, length: 10}, []Item{}}
	items := []Item{{Box{width: 2, height: 5, length: 4}}, {Box{width: 2, height: 5, length: 4}}, {Box{width: 2, height: 5, length: 4}}, {Box{width: 2, height: 5, length: 4}}}
	pool := NewPool(items)

	containers := breakSpace(container) // serial

	// here ends the serial section, from now own everything can be in parallel, besides the last segment

	for _, c := range containers {
		go c.pickItems(pool)
	}

	fmt.Println("Hello, 世界")
}

func breakSpace(full Container) []Container {
	// todo: break the space
	return []Container{
		{Box{height: 10, width: 5, length: 10}, []Item{}},
		{Box{height: 10, width: 5, length: 10}, []Item{}},
	}
}
