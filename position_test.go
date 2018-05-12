package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Position_NewPos_Successful(t *testing.T) {
	x := 1.0
	y := 2.0

	pos := NewPosition(x, y)

	assert.Equal(t, x, pos.X)
	assert.Equal(t, y, pos.Y)
}
