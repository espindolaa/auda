package io

import (
	"encoding/json"
	"io/ioutil"
	"tcc/model"
)

type Visual struct {
	Boxes model.Utilization
	Areas []model.Container
}

func GenerateJson(solution model.Utilization, containers []model.Container, filename string) {
	j := Visual{Boxes: solution, Areas: containers}

	file, _ := json.MarshalIndent(j, "", " ")

	_ = ioutil.WriteFile(filename, file, 0644)
}
