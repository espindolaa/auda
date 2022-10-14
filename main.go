package auda

import (
	"tcc/model"
)

func isValidSolution(containers []model.Container) bool {
	fullSpace := model.Utilization{}
	for _, c := range containers {
		if !fullSpace.Append(c.Sorted) {
			return false
		}
	}

	return true
}

func waitForCompletion(ch chan model.Utilization, numberOfContainers int) []model.Utilization {
	utilizations := []model.Utilization{}

	for {
		v := <-ch
		utilizations = append(utilizations, v)

		if len(utilizations) == numberOfContainers {
			break
		}

	}

	return utilizations
}

func main() {
	// todo: move to cmd input
	container := model.NewContainer(model.NewBox(10, 10, 10), model.Point{X: 0, Y: 0, Z: 0})

	items := []model.Item{model.NewItem(2, 5, 4), model.NewItem(2, 5, 4), model.NewItem(2, 5, 4), model.NewItem(2, 5, 4)}
	pool := model.NewPool(items)

	containers := container.BreakSpace(container) // serial

	channel := make(chan model.Utilization)

	for _, c := range containers {
		go c.Sort(pool, channel)
	}

	waitForCompletion(channel, len(containers))

	if isValidSolution(containers) {
		return
	}
}
