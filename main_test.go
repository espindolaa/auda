package auda

import (
	"tcc/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SomeSolutionIsFound(t *testing.T) {
	container := model.NewContainer(model.NewBox(10, 10, 10), model.Point{X: 0, Y: 0, Z: 0})
	items := []model.Item{model.NewItem(2, 5, 4), model.NewItem(2, 5, 4), model.NewItem(2, 5, 4), model.NewItem(2, 5, 4)}

	found, result := core(container, items)

	assert.True(t, found)
	assert.Equal(t, len(result), 1)
}
