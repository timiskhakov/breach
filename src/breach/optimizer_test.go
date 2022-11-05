package breach

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge_Case1_OneElementMatches(t *testing.T) {
	seq1 := []byte{0x55, 0x1C}
	seq2 := []byte{0x1C, 0x55, 0x1C}

	actual := merge(seq1, seq2)

	assert.Equal(t, []byte{0x55, 0x1C, 0x55, 0x1C}, actual)
}

func TestMerge_Case1_TwoElementsMatch(t *testing.T) {
	seq1 := []byte{0x55, 0x1C}
	seq2 := []byte{0x55, 0x1C, 0x1C}

	actual := merge(seq1, seq2)

	assert.Equal(t, []byte{0x55, 0x1C, 0x1C}, actual)
}

func TestMerge_Case1_NoMatch(t *testing.T) {
	seq1 := []byte{0x55, 0xFF}
	seq2 := []byte{0x55, 0x1C, 0x1C}

	actual := merge(seq1, seq2)

	assert.Equal(t, []byte{0x55, 0xFF, 0x55, 0x1C, 0x1C}, actual)
}