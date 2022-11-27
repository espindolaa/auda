package model

import (
	"github.com/gedex/bp3d"
)

type PackingAdapter interface {
	Pack(Container) Utilization
}

type BP3DAdapter struct {
}

func FromItem(i Item) *bp3d.Item {
	return bp3d.NewItem(i.Label, i.Width, i.Height, i.Length, i.Weight)
}

func ToItem(i *bp3d.Item) Item {
	return NewItem(i.Name, i.Width, i.Depth, i.Height, i.Weight)
}

func ToPosition(i *bp3d.Item, bottomLeft Point) Position {
	dimensions := i.GetDimension()
	item := NewItem(i.Name, dimensions[0], dimensions[2], dimensions[1], i.Weight)
	position := NewPosition(item, Point{i.Position[0] + bottomLeft.X, i.Position[2] + bottomLeft.Y, i.Position[1] + bottomLeft.Z})
	return position
}

func (b3pd BP3DAdapter) Pack(c Container) Utilization {
	packer := bp3d.NewPacker()

	packer.AddBin(bp3d.NewBin("", c.Width, c.Height, c.Length, c.MaxWeight))

	m := make(map[string]Item)

	for _, item := range c.unsortedItems {
		m[item.Label] = item
		packer.AddItem(FromItem(item))
	}

	packer.Pack()

	var result Utilization

	for _, i := range packer.Bins[0].Items {
		result.Add(ToPosition(i, c.LeftBottomCorner))
	}

	for _, i := range packer.UnfitItems {
		result.AddUnused(m[i.Name])
	}

	return result
}
