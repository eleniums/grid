package grid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Grid_NewGrid_Successful(t *testing.T) {
	grid := NewGrid()

	assert.NotNil(t, grid)
	assert.Empty(t, grid)
}

func Test_Grid_Add_Successful(t *testing.T) {
	grid := NewGrid()

	x := 1.0
	y := 2.0
	expected := "test"

	grid.Add(x, y, expected)

	assert.Equal(t, expected, grid[NewPosition(x, y)])
}

func Test_Grid_Add_Overwrite(t *testing.T) {
	grid := NewGrid()

	x := 1.0
	y := 2.0
	expected := "test2"

	grid.Add(x, y, "test1")
	grid.Add(x, y, expected)

	assert.Equal(t, expected, grid[NewPosition(x, y)])
}

func Test_Grid_Retrieve_Successful(t *testing.T) {
	grid := NewGrid()

	x := 1.0
	y := 2.0
	expected := "test"

	grid.Add(x, y, expected)

	value, ok := grid.Retrieve(x, y)

	assert.True(t, ok)
	assert.Equal(t, expected, value)
}

func Test_Grid_Retrieve_NoValue(t *testing.T) {
	grid := NewGrid()

	x := 1.0
	y := 2.0

	value, ok := grid.Retrieve(x, y)

	assert.False(t, ok)
	assert.Nil(t, value)
}

func Test_Grid_Delete_Successful(t *testing.T) {
	grid := NewGrid()

	x := 1.0
	y := 2.0

	grid.Add(x, y, "test")

	grid.Delete(x, y)

	value, ok := grid[NewPosition(x, y)]
	assert.False(t, ok)
	assert.Nil(t, value)
}

func Test_Grid_Delete_NoValue(t *testing.T) {
	grid := NewGrid()

	x := 1.0
	y := 2.0

	grid.Delete(x, y)

	value, ok := grid[NewPosition(x, y)]
	assert.False(t, ok)
	assert.Nil(t, value)
}
