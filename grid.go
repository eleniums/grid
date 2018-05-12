package grid

// Grid is a grid of values of any type.
type Grid map[Position]interface{}

// NewGrid will create a new empty grid.
func NewGrid() Grid {
	return Grid{}
}

// Add a value at position (x, y).
func (g Grid) Add(x, y float64, value interface{}) {
	g[NewPosition(x, y)] = value
}
