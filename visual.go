package auda

import (
	"encoding/json"
	"io/ioutil"
	"tcc/model"
)

type Visual struct {
	Boxes model.Utilization
	Areas []model.Container
}

func generateJson(solution model.Utilization, containers []model.Container) {
	j := Visual{Boxes: solution, Areas: containers}

	file, _ := json.MarshalIndent(j, "", " ")

	_ = ioutil.WriteFile("result.json", file, 0644)
}
