package auda

import (
	"fmt"
	"tcc/model"
)

func formFinalSolution(solution []model.Utilization) model.Utilization {
	fullSpace := model.Utilization{}
	fmt.Printf("Containers usados: %d", len(solution))
	for _, u := range solution {
		fullSpace.Append(u)
	}

	return fullSpace
}

func waitForCompletion(ch chan model.Utilization, numberOfContainers int) []model.Utilization {
	utilizations := []model.Utilization{}

	for v := range ch {
		utilizations = append(utilizations, v)

		if len(utilizations) == numberOfContainers {
			break
		}

	}

	return utilizations
}

func core(container model.Container, items []model.Item) model.Utilization {
	pool := model.NewPool(items)

	containers := model.BreakSpace([]model.Container{container}, len(items))

	channel := make(chan model.Utilization)

	for _, c := range containers {
		go c.Sort(pool, channel)
	}

	pool.AllowTake()

	solutions := waitForCompletion(channel, len(containers))
	result := formFinalSolution(solutions)

	generateJson(result, containers)

	return result
}
