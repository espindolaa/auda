package auda

import (
	"fmt"
	"testing"

	"github.com/gedex/bp3d"
)

// run with go test -bench=. -benchtime=1x

const (
	testFile = "tests/thpack1.txt"
)

func BenchmarkHeuristic(b *testing.B) {
	b.StopTimer()
	tests := ParseTestFile(testFile)

	fmt.Println("Heuristic results")

	b.SetParallelism(1)
	b.StartTimer()

	for _, t := range tests {
		result := core(t.Container, t.Items)

		fmt.Printf("Test ID: %s \n", t.ID)
		fmt.Printf("Volume occupied: %f \n", result.Volume)
		fmt.Printf("Number of items: %d \n", len(result.Allocated))
		fmt.Println()
	}
	b.Name()
}

func BenchmarkSimple(b *testing.B) {
	b.StopTimer()
	tests := ParseTestFile(testFile)

	type SimpleTest struct {
		ID        string
		Container *bp3d.Bin
		Items     []*bp3d.Item
	}

	preparedTests := []SimpleTest{}
	for _, t := range tests {
		items := []*bp3d.Item{}
		for _, i := range t.Items {
			items = append(items, bp3d.NewItem(i.Label, i.Width, i.Height, i.Length, i.Weight))
		}

		preparedTests = append(preparedTests, SimpleTest{
			ID:        t.ID,
			Container: bp3d.NewBin("Container", t.Container.Width, t.Container.Height, t.Container.Length, 1000),
			Items:     items,
		})
	}

	fmt.Println("Simple results")

	b.SetParallelism(1)
	b.StartTimer()

	for _, t := range preparedTests {
		p := bp3d.NewPacker()

		p.AddBin(t.Container)

		for _, i := range t.Items {
			p.AddItem(i)
		}

		p.Pack()

		volume := 0.0

		for _, i := range p.Bins[0].Items {
			volume += i.Depth * i.Height * i.Width
		}

		fmt.Printf("Test ID: %s \n", t.ID)
		fmt.Printf("Volume occupied: %f \n", volume)
		fmt.Printf("Number of items: %d \n", len(p.Bins[0].Items))
		fmt.Println()
	}
}
