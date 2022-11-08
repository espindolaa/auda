package model

type Point struct {
	X float64
	Y float64
	Z float64
}

type box struct {
	Width  float64
	Length float64
	Height float64
}

func NewBox(width, length, height float64) box {
	return box{Width: width, Length: length, Height: height}
}

func (b box) Volume() float64 {
	return b.Height * b.Width * b.Length
}

type Item struct {
	box
	Weight float64
	Label  string
}

func NewItem(label string, width, length, height, weight float64) Item {
	return Item{box: box{Width: width, Length: length, Height: height}, Weight: weight, Label: label}
}

type Container struct {
	box
	MaxWeight        float64
	LeftBottomCorner Point
	unsortedItems    []Item
	breakingStrategy BreakingStrategy
	packingAlgorithm PackingAdapter
	Sorted           Utilization
}

func NewContainer(dimensions box, maxWeight float64, position Point, packingAlgorithm PackingAdapter) Container {
	return Container{dimensions, maxWeight, position, []Item{}, &ByColumn{}, packingAlgorithm, Utilization{}}
}

func (c Container) Sort(pool *Pool, channel chan Utilization) {
	c.PickItems(pool)
	c.Sorted = c.packingAlgorithm.Pack(c)

	channel <- c.Sorted
}

func BreakSpace(spaces []Container, numberOfItems int) []Container {
	if numberOfItems > 200 { // magic number
		// call another break
		numberOfItems = numberOfItems / 2
		brokenSpaces := []Container{}

		for _, s := range spaces {
			broken := s.breakingStrategy.Break(s)
			brokenSpaces = append(brokenSpaces, broken...)
		}

		return BreakSpace(brokenSpaces, numberOfItems)
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

type Utilization struct {
	Volume    float64
	Allocated []Position
	Unused    []Item
}

func (u *Utilization) Add(position Position) {
	u.Allocated = append(u.Allocated, position)
	u.Volume += position.Item.box.Volume()
}

func (u *Utilization) AddUnused(item Item) {
	u.Unused = append(u.Unused, item)
}

func (u *Utilization) Append(ut Utilization) {
	u.Unused = append(u.Unused, ut.Unused...)
	u.Allocated = append(u.Allocated, ut.Allocated...)
	u.Volume += ut.Volume
}

type Position struct {
	Item             Item
	LeftBottomCorner Point
}

func NewPosition(item Item, leftBottomCorner Point) Position {
	return Position{
		Item:             item,
		LeftBottomCorner: leftBottomCorner,
	}
}
