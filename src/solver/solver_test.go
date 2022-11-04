package solver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve_Case1(t *testing.T) {
	matrix := [][]byte{
		{0xE9, 0xE9, 0x7A, 0xBD, 0x55, 0x55},
		{0x1C, 0x1C, 0x1C, 0x7A, 0x55, 0xE9},
		{0x1C, 0x7A, 0x7A, 0x1C, 0x55, 0x1C},
		{0xBD, 0xE9, 0x55, 0x7A, 0x55, 0x7A},
		{0x55, 0x55, 0x55, 0x7A, 0x55, 0x1C},
		{0xBD, 0xBD, 0xE9, 0x1C, 0x55, 0xE9},
	}

	actual := Solve(matrix, []byte{0x55, 0x55, 0x7A})

	assert.Equal(t, []Point{{0, 4}, {1, 4}, {1, 3}}, actual)
}

func TestSolve_Case2(t *testing.T) {
	matrix := [][]byte{
		{0xBD, 0xE9, 0x1C, 0xBD, 0xBD},
		{0x1C, 0x1C, 0xBD, 0xBD, 0x55},
		{0xBD, 0xBD, 0x55, 0xBD, 0xBD},
		{0x1C, 0x55, 0x55, 0xE9, 0xBD},
		{0x55, 0x1C, 0xBD, 0x55, 0x55},
	}

	actual := Solve(matrix, []byte{0x55, 0x1C})

	assert.Equal(t, []Point{{1, 4}, {1, 0}}, actual)
}
