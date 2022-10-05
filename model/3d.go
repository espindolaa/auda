package model

type Point struct {
	X int
	Y int
	Z int
}

type box struct {
	Width  int
	Length int
	Height int
}

func NewBox(width, length, height int) box {
	return box{Width: width, Length: length, Height: height}
}

func (b box) Volume() int {
	return b.Height * b.Width * b.Length
}

type Item struct {
	box
}

func NewItem(width, length, height int) Item {
	return Item{box{Width: width, Length: length, Height: height}}
}

type Container struct {
	box
	LeftBottomPoint  Point
	unsortedItems    []Item
	breakingStrategy BreakingStrategy
	Sorted           Utilization
}

func NewContainer(dimensions box, position Point) Container {
	return Container{dimensions, position, []Item{}, &ByColumn{}, Utilization{}}
}

func (c *Container) BreakSpace(full Container) []Container {
	spaces := c.breakingStrategy.Break(*c)

	for _, space := range spaces {
		if space.Volume() > 300 { // extract magic number
			broken := space.BreakSpace(space)
			spaces = spaces[1:]
			spaces = append(spaces, broken...)
		}
	}

	return spaces
}

func (c *Container) PickItems(pool *Pool) {
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

func (c *Container) ArrangeItems() {
	// todo: call the packing algorithm using c.unsortedItems
	c.Sorted = Utilization{}
}

type Utilization struct {
	Volume    int
	Allocated []Position
	Unused    []Item
}

func (u *Utilization) Append(ut Utilization) bool { // o (n*m)
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
