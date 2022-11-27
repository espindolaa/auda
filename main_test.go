package auda

import (
	"fmt"
	"tcc/model"
	"testing"

	"github.com/gedex/bp3d"
)

// go test -run=. -bench=. -benchtime=1ns

const (
	testFile = "tests/thpack2.txt"
)

func BenchmarkHeuristic(b *testing.B) {
	b.StopTimer()
	tests := ParseTestFile(testFile)

	fmt.Println("Heuristic results")

	b.SetParallelism(1)
	b.StartTimer()

	totalVolume := 0
	totalItems := 0

	for _, t := range tests {
		result := heuristic(t.Container, t.Items)

		fmt.Printf("Test ID: %s \n", t.ID)
		fmt.Printf("%f \n", result.Volume)
		fmt.Printf("%d \n", len(result.Allocated))

		totalVolume += int(result.Volume)
		totalItems += len(result.Allocated)
	}
	b.StopTimer()

	fmt.Printf("Total volume: %d \n", totalVolume)
	fmt.Printf("Total items: %d \n", totalItems)
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
			items = append(items, model.FromItem(i))
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

	totalVolume := 0
	totalItems := 0

	for _, t := range preparedTests {
		bin := simple(t.Container, t.Items)

		volume := 0.0

		for _, i := range bin.Items {
			volume += i.Depth * i.Height * i.Width
		}

		fmt.Printf("Test ID: %s \n", t.ID)
		fmt.Printf("%f \n", volume)
		fmt.Printf("%d \n", len(bin.Items))

		totalVolume += int(volume)
		totalItems += len(bin.Items)
	}

	b.StopTimer()
	fmt.Printf("Total volume: %d \n", totalVolume)
	fmt.Printf("Total items: %d \n", totalItems)
}
