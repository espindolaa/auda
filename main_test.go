package auda

import (
	"fmt"
	"log"
	"tcc/model"
	"testing"

	"github.com/gedex/bp3d"
	"github.com/stretchr/testify/assert"
)

func GetLargeShipment() []model.Item {
	return []model.Item{
		model.NewItem("Geladeira 1", 60, 70, 180, 60),
		model.NewItem("Geladeira 2", 60, 70, 180, 60),
		model.NewItem("Geladeira 3", 60, 70, 180, 60),
		model.NewItem("Geladeira 4", 60, 70, 180, 60),
		model.NewItem("Geladeira 5", 60, 70, 180, 60),
		model.NewItem("Geladeira 6", 60, 70, 180, 60),
		model.NewItem("Geladeira 7", 60, 70, 180, 60),
		model.NewItem("Geladeira 8", 60, 70, 180, 60),
		model.NewItem("Geladeira 9", 60, 70, 180, 60),
		// model.NewItem("Geladeira 10", 60, 70, 180, 60),
		// model.NewItem("Geladeira 11", 60, 70, 180, 60),
		// model.NewItem("Geladeira 12", 60, 70, 180, 60),
		// model.NewItem("Geladeira 13", 60, 70, 180, 60),
		// model.NewItem("Geladeira 14", 60, 70, 180, 60),
		// model.NewItem("Geladeira 15", 60, 70, 180, 60),
		// model.NewItem("Geladeira 16", 60, 70, 180, 60),
		// model.NewItem("Geladeira 17", 60, 70, 180, 60),
		// model.NewItem("Geladeira 18", 60, 70, 180, 60),
		// model.NewItem("Geladeira 19", 60, 70, 180, 60),
		// model.NewItem("Geladeira 20", 60, 70, 180, 60),
		// model.NewItem("Fogão 1", 60, 70, 70, 45),
		// model.NewItem("Fogão 2", 60, 70, 70, 45),
		// model.NewItem("Fogão 3", 60, 70, 70, 45),
		// model.NewItem("Fogão 4", 60, 70, 70, 45),
		// model.NewItem("Fogão 5", 60, 70, 70, 45),
		// model.NewItem("Fogão 6", 60, 70, 70, 45),
		// model.NewItem("Fogão 7", 60, 70, 70, 45),
		// model.NewItem("Fogão 8", 60, 70, 70, 45),
		// model.NewItem("Fogão 9", 60, 70, 70, 45),
		// model.NewItem("Fogão 10", 60, 70, 70, 45),
		// model.NewItem("Sofá 1", 150, 100, 90, 50),
		// model.NewItem("Sofá 2", 150, 100, 90, 50),
		// model.NewItem("Sofá 3", 150, 100, 90, 50),
		// model.NewItem("Sofá 4", 150, 100, 90, 50),
		// model.NewItem("Sofá 5", 150, 100, 90, 50),
		// model.NewItem("Sofá 6", 150, 100, 90, 50),
		// model.NewItem("Sofá 7", 150, 100, 90, 50),
		// model.NewItem("Sofá 8", 150, 100, 90, 50),
		// model.NewItem("Sofá 9", 150, 100, 90, 50),
		// model.NewItem("Sofá 10", 150, 100, 90, 50),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
		model.NewItem("Microondas 1", 45, 40, 28, 15),
		model.NewItem("Microondas 2", 45, 40, 28, 15),
		model.NewItem("Microondas 3", 45, 40, 28, 15),
		model.NewItem("Microondas 4", 45, 40, 28, 15),
		model.NewItem("Microondas 5", 45, 40, 28, 15),
		model.NewItem("Microondas 6", 45, 40, 28, 15),
		model.NewItem("Microondas 7", 45, 40, 28, 15),
		model.NewItem("Microondas 8", 45, 40, 28, 15),
		model.NewItem("Microondas 9", 45, 40, 28, 15),
		model.NewItem("Microondas 10", 45, 40, 28, 15),
	}
}

func BenchmarkHeuristic(b *testing.B) {
	container := model.NewContainer(model.NewBox(244, 1220, 227), 10000, model.Point{X: 0, Y: 0, Z: 0}, model.B3PDAdapter{})
	items := GetLargeShipment()

	result := core(container, items)

	fmt.Printf("Volume benchmark heuristic: %f", result.Volume)
	fmt.Printf("Number of items: %d", len(result.Allocated))
}

func BenchmarkSimple(b *testing.B) {
	p := bp3d.NewPacker()

	// Add bins.
	p.AddBin(bp3d.NewBin("Container", 244, 1220, 227, 1000))

	items := GetLargeShipment()

	for _, i := range items {
		p.AddItem(bp3d.NewItem(i.Label, i.Width, i.Height, i.Length, i.Weight))
	}

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	volume := 0.0

	for _, i := range p.Bins[0].Items {
		volume += i.Depth * i.Height * i.Width
	}

	fmt.Printf("Volume benchmark simple: %f", volume)
	fmt.Printf("Number of items: %d", len(p.Bins[0].Items))
}

func Test_SomeSolutionIsFound(t *testing.T) {
	container := model.NewContainer(model.NewBox(244, 226, 600), 10000, model.Point{X: 0, Y: 0, Z: 0}, model.B3PDAdapter{})
	items := []model.Item{model.NewItem("1", 2, 5, 4, 1), model.NewItem("2", 2, 5, 4, 2), model.NewItem("3", 2, 5, 4, 3), model.NewItem("4", 2, 5, 4, 4)}

	result := core(container, items)

	assert.NotNil(t, result)
}

func Test_CallPackingAlgorithm(t *testing.T) {
	p := bp3d.NewPacker()

	// Add bins.
	p.AddBin(bp3d.NewBin("Container", 244, 226, 600, 1000))

	items := []model.Item{model.NewItem("1", 2, 5, 4, 1), model.NewItem("2", 2, 5, 4, 2), model.NewItem("3", 2, 5, 4, 3), model.NewItem("4", 2, 5, 4, 4)}

	for _, i := range items {
		p.AddItem(bp3d.NewItem(i.Label, i.Width, i.Height, i.Length, i.Weight))
	}

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	// Each bin, b, in p.Bins might have packed items in b.Items
}
