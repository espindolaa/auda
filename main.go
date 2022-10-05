package auda

import "tcc/model"

func validateSolution(containers []model.Container) bool {
	fullSpace := model.Utilization{}
	for _, c := range containers {
		if !fullSpace.Append(c.Sorted) {
			return false
		}
	}

	return true
}

func main() {
	// todo: move to cmd input
	container := model.NewContainer(model.NewBox(10, 10, 10), model.Point{X: 0, Y: 0, Z: 0})

	items := []model.Item{model.NewItem(2, 5, 4), model.NewItem(2, 5, 4), model.NewItem(2, 5, 4), model.NewItem(2, 5, 4)}
	pool := model.NewPool(items)

	containers := container.BreakSpace(container) // serial

	// here ends the serial section, from now own everything can be in parallel, besides the last segment

	for _, c := range containers {
		go c.PickItems(pool)
		go c.ArrangeItems()
	}

	// last segment

	validateSolution(containers)
}
