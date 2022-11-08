package auda

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tcc/model"
	"testing"

	"github.com/gedex/bp3d"
)

// run with go test -bench=. -benchtime=1x

func GetCopiesOfItems(amount int, name string, length, width, height, weight float64) []model.Item {
	items := []model.Item{}

	for i := 1; i <= amount; i++ {
		n := fmt.Sprintf("%s - %d", name, i)
		items = append(items, model.NewItem(n, width, length, height, weight))
	}

	return items
}

type TestInstance struct {
	ID        string
	Container model.Container
	Items     []model.Item
}

func ParseTestFile(name string) []TestInstance {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tests := []TestInstance{}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for {
		i := 0
		test := TestInstance{}
		test.ID = lines[i]
		i++
		c := strings.Split(lines[i], " ")
		containerL, _ := strconv.ParseFloat(c[1], 64)
		containerW, _ := strconv.ParseFloat(c[0], 64)
		containerH, _ := strconv.ParseFloat(c[2], 64)
		test.Container = model.NewContainer(model.NewBox(containerL, containerW, containerH), 10000, model.Point{X: 0, Y: 0, Z: 0}, model.B3PDAdapter{})
		i++
		n, _ := strconv.Atoi(lines[i])
		i++
		items := []model.Item{}
		for j := 0; j < n; j++ {
			l := strings.Split(lines[i], " ")
			itemL, _ := strconv.ParseFloat(l[3], 64)
			itemW, _ := strconv.ParseFloat(l[1], 64)
			itemH, _ := strconv.ParseFloat(l[5], 64)
			copies, _ := strconv.Atoi(l[7])
			items = append(items, GetCopiesOfItems(copies, strconv.Itoa(n), itemL, itemW, itemH, 0)...)
			test.Items = append(test.Items, items...)
			i++
		}
		tests = append(tests, test)
		lines = lines[i:]
		if len(lines) == 0 {
			break
		}
	}

	return tests
}

func BenchmarkHeuristic(b *testing.B) {
	b.StopTimer()
	tests := ParseTestFile("test.txt")

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
}

func BenchmarkSimple(b *testing.B) {
	b.StopTimer()
	tests := ParseTestFile("test.txt")

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
