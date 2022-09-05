package auda

import "fmt"

type Box struct {
	width  int
	length int
	height int
}

type Container struct {
	Box
}

type Item struct {
	Box
}

func (b Box) Volume() int {
	return b.height * b.width * b.length
}

func main() {
	// todo: move to cmd input
	container := Container{Box{height: 10, width: 10, length: 10}}
	containers := breakSpace(container)

	fmt.Println("Hello, 世界")
}

func breakSpace(full Container) []Container {
	// todo: break the space
	return []Container{Container{Box{height: 10, width: 10, length: 5}}, Container{Box{height: 10, width: 10, length: 5}}}
}
