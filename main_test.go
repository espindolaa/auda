package auda

import (
	"fmt"
	"syscall"
	"tcc/model"
	"testing"

	"github.com/gedex/bp3d"
)

// run with go test -bench=. -benchtime=1x

const (
	testFile = "tests/thpack1.txt"
)

func GetCPU() int64 {
	usage := new(syscall.Rusage)
	syscall.Getrusage(syscall.RUSAGE_SELF, usage)
	return usage.Utime.Nano() + usage.Stime.Nano()
}

func BenchmarkHeuristic(b *testing.B) {
	b.StopTimer()
	tests := ParseTestFile(testFile)

	fmt.Println("Heuristic results")

	b.SetParallelism(1)
	b.StartTimer()

	for _, t := range tests {
		result := heuristic(t.Container, t.Items)

		fmt.Printf("Test ID: %s \n", t.ID)
		fmt.Printf("Volume occupied: %f \n", result.Volume)
		fmt.Printf("Number of items: %d \n", len(result.Allocated))
		fmt.Println()
	}
	b.StopTimer()
	fmt.Printf("CPU Time: %d \n", GetCPU())
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

	for _, t := range preparedTests {
		bin := simple(t.Container, t.Items)

		volume := 0.0

		for _, i := range bin.Items {
			volume += i.Depth * i.Height * i.Width
		}

		fmt.Printf("Test ID: %s \n", t.ID)
		fmt.Printf("Volume occupied: %f \n", volume)
		fmt.Printf("Number of items: %d \n", len(bin.Items))
		fmt.Println()
	}

	b.StopTimer()
	fmt.Printf("CPU Time: %d \n", GetCPU())
}
