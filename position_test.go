package grid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Position_NewPos_Success(t *testing.T) {
	x := 1.0
	y := 2.0

	pos := NewPosition(x, y)

	assert.Equal(t, x, pos.X)
	assert.Equal(t, y, pos.Y)
}
