package auda

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tcc/model"
)

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

	i := 0
	numberOfProblems, _ := strconv.Atoi(lines[i])
	i++

	for p := 0; p < numberOfProblems; p++ {
		test := TestInstance{}
		id := strings.Split(lines[i], " ")
		test.ID = id[0]
		i++
		c := strings.Split(lines[i], " ")
		containerL, _ := strconv.ParseFloat(c[1], 64)
		containerW, _ := strconv.ParseFloat(c[0], 64)
		containerH, _ := strconv.ParseFloat(c[2], 64)
		test.Container = model.NewContainer(model.NewBox(containerL, containerW, containerH), 10000, model.Point{X: 0, Y: 0, Z: 0}, model.BP3DAdapter{})
		i++
		n, _ := strconv.Atoi(lines[i])
		i++
		for j := 0; j < n; j++ {
			l := strings.Split(lines[i], " ")
			itemL, _ := strconv.ParseFloat(l[3], 64)
			itemW, _ := strconv.ParseFloat(l[1], 64)
			itemH, _ := strconv.ParseFloat(l[5], 64)
			copies, _ := strconv.Atoi(l[7])
			copied := GetCopiesOfItems(copies, strconv.Itoa(n), itemL, itemW, itemH, 0)
			test.Items = append(test.Items, copied...)
			i++
		}
		tests = append(tests, test)
	}

	return tests
}
