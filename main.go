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

func core(container model.Container, items []model.Item) (bool, []model.Utilization) {
	pool := model.NewPool(items)

	containers := container.BreakSpace(container) // serial

	channel := make(chan model.Utilization)

	for _, c := range containers {
		go c.Sort(pool, channel)
	}

	solution := waitForCompletion(channel, len(containers))

	if isValidSolution(containers) {
		return false, nil
	}

	return true, solution
}
