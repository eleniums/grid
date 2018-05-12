package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Grid_NewGrid_Successful(t *testing.T) {
	grid := NewGrid()

	assert.NotNil(t, grid)
	assert.Empty(t, grid)
}

func Test_Grid_Add_Successful(t *testing.T) {
	grid := NewGrid()

	expected := "test"

	grid.Add(1, 1, expected)

	assert.Equal(t, expected, grid[NewPosition(1, 1)])
}
