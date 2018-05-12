package grid

// Position is a 2D point in space.
type Position struct {
	X float64
	Y float64
}

// NewPosition will create a new position.
func NewPosition(x, y float64) Position {
	return Position{
		X: x,
		Y: y,
	}
}
