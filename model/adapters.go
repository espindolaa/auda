package model

import (
	"log"

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

func ToUtilization(bin *bp3d.Bin, unfit []*bp3d.Item) Utilization {
	var result Utilization

	for _, i := range bin.Items {
		item := NewItem(i.Name, i.Width, i.Depth, i.Height, i.Weight)
		containerPoint := Point{0, 0, 0}
		position := NewPosition(item, Point{i.Position[0] + containerPoint.X, i.Position[2] + containerPoint.Y, i.Position[1] + containerPoint.Z})
		result.Add(position)
	}

	for _, i := range unfit {
		result.AddUnused(NewItem(i.Name, i.Width, i.Depth, i.Height, i.Weight))
	}

	return result
}

func (b3pd BP3DAdapter) Pack(c Container) Utilization {
	packer := bp3d.NewPacker()

	packer.AddBin(bp3d.NewBin("", c.Width, c.Height, c.Length, c.MaxWeight))

	m := make(map[string]Item)

	for _, item := range c.unsortedItems {
		m[item.Label] = item
		packer.AddItem(FromItem(item))
	}

	if err := packer.Pack(); err != nil { // todo: handle error
		log.Fatal(err)
	}

	var result Utilization

	for _, i := range packer.Bins[0].Items {
		item := m[i.Name]
		containerPoint := c.LeftBottomCorner
		position := NewPosition(item, Point{i.Position[0] + containerPoint.X, i.Position[2] + containerPoint.Y, i.Position[1] + containerPoint.Z})
		result.Add(position)
	}

	for _, i := range packer.UnfitItems {
		result.AddUnused(m[i.Name])
	}

	return result
}
