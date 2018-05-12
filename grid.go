package grid

// Grid is a grid of values of any type.
type Grid map[Position]interface{}

// NewGrid will create a new empty grid.
func NewGrid() Grid {
	return Grid{}
}

// Add value at position (x, y).
func (g Grid) Add(x, y float64, value interface{}) {
	g[NewPosition(x, y)] = value
}

// Retrieve value at position (x, y).
func (g Grid) Retrieve(x, y float64) (interface{}, bool) {
	value, ok := g[NewPosition(x, y)]
	return value, ok
}

// Delete value at position (x, y).
func (g Grid) Delete(x, y float64) {
	delete(g, NewPosition(x, y))
}
