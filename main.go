package auda

import "sync"

type BreakingStrategy interface {
	Break(full Container) []Container
}

type ByColumn struct {
	BreakingStrategy
}

func (c *ByColumn) Break(full Container) []Container {
	half := full.length / 2
	f := full
	f.length = half

	s := full
	s.length = half
	s.leftBottomPoint = Point{x: full.leftBottomPoint.x, y: full.leftBottomPoint.y + half, z: full.leftBottomPoint.z}

	return []Container{f, s}
}

type Box struct {
	width  int
	length int
	height int
}

type Point struct {
	x int
	y int
	z int
}

type Container struct {
	Box
	leftBottomPoint  Point
	unsortedItems    []Item
	breakingStrategy BreakingStrategy
	sorted           Utilization
}

func NewContainer(dimensions Box, position Point) Container {
	return Container{dimensions, position, []Item{}, &ByColumn{}, Utilization{}}
}

func (c *Container) breakSpace(full Container) []Container {
	spaces := c.breakingStrategy.Break(*c)

	for _, space := range spaces {
		if space.Volume() > 300 { // extract magic number
			broken := space.breakSpace(space)
			spaces = spaces[1:]
			spaces = append(spaces, broken...)
		}
	}

	return spaces
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

type Utilization struct {
	Volume    int
	Allocated []Position
	Unused    []Item
}

func (u *Utilization) append(ut Utilization) bool { // o (n*m)
	for _, a := range u.Allocated {
		for _, al := range ut.Allocated {
			if conflict(a, al) {
				return false
			}
		}
	}

	u.Unused = append(u.Unused, ut.Unused...)
	u.Allocated = append(u.Allocated, ut.Allocated...)
	u.Volume += ut.Volume

	return true
}

func conflict(first, second Position) bool {
	// todo: guarantee that the two boxes don't overlap (1D cut)
	return false
}

type Position struct {
	item             Item
	leftBottomCorner Point
}

func (c *Container) arrangeItems() {
	// todo: call the packing algorithm using c.unsortedItems
	c.sorted = Utilization{}
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

func validateSolution(containers []Container) bool {
	fullSpace := Utilization{}
	for _, c := range containers {
		if !fullSpace.append(c.sorted) {
			return false
		}
	}

	return true
}

func main() {
	// todo: move to cmd input
	container := NewContainer(Box{height: 10, width: 10, length: 10}, Point{x: 0, y: 0, z: 0})

	items := []Item{{Box{width: 2, height: 5, length: 4}}, {Box{width: 2, height: 5, length: 4}}, {Box{width: 2, height: 5, length: 4}}, {Box{width: 2, height: 5, length: 4}}}
	pool := NewPool(items)

	containers := container.breakSpace(container) // serial

	// here ends the serial section, from now own everything can be in parallel, besides the last segment

	for _, c := range containers {
		go c.pickItems(pool)
		go c.arrangeItems()
	}

	// last segment

	validateSolution(containers)
}
