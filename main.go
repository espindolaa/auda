package auda

import (
	"tcc/model"

	"github.com/gedex/bp3d"
)

func formFinalSolution(solution []model.Utilization) model.Utilization {
	fullSpace := model.Utilization{}
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

func heuristic(container model.Container, items []model.Item) model.Utilization {
	pool := model.NewPool(items)

	containers := model.BreakSpace([]model.Container{container}, len(items))

	channel := make(chan model.Utilization)

	for _, c := range containers {
		go c.Sort(pool, channel)
	}

	pool.AllowTake()

	solutions := waitForCompletion(channel, len(containers))
	result := formFinalSolution(solutions)

	// io.GenerateJson(result, containers, "heuristic.json")

	return result
}

func simple(container *bp3d.Bin, items []*bp3d.Item) *bp3d.Bin {
	p := bp3d.NewPacker()

	p.AddBin(container)

	for _, i := range items {
		p.AddItem(i)
	}

	p.Pack()

	// result := model.ToUtilization(p.Bins[0], p.UnfitItems)
	// io.GenerateJson(result, []model.Container{
	// 	model.NewContainer(model.NewBox(container.Width, container.Depth, container.Height), 10000, model.Point{X: 0, Y: 0, Z: 0}, nil),
	// }, "simple.json")

	return p.Bins[0]
}
