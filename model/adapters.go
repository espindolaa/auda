package model

import (
	"log"

	"github.com/gedex/bp3d"
)

type PackingAdapter interface {
	Pack(Container) Utilization
}

type B3PDAdapter struct {
}

func (b3pd B3PDAdapter) Pack(c Container) Utilization {
	packer := bp3d.NewPacker()

	packer.AddBin(bp3d.NewBin("", c.Width, c.Height, c.Length, c.MaxWeight))

	m := make(map[string]Item)

	for _, item := range c.unsortedItems {
		m[item.Label] = item
		packer.AddItem(bp3d.NewItem(item.Label, item.Width, item.Height, item.Length, item.Weight))
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