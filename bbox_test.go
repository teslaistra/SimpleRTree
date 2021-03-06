package SimpleRTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBBox_Extend(t *testing.T) {
	testCases := []struct {
		b1, b2   rBBox
		expected rBBox
	}{
		{
			b1:       rBBox{0, 0, 1, 1},
			b2:       rBBox{1, 1, 2, 2},
			expected: rBBox{0, 0, 2, 2},
		},
	}

	for _, tc := range testCases {
		aB1 := bbox2VectorBBox(tc.b1)
		aB2 := bbox2VectorBBox(tc.b2)
		result := vectorBBoxExtend(aB1, aB2)
		assert.Equal(t, tc.expected, result.toBBox())
		assert.Equal(t, tc.expected, tc.b1.extend(tc.b2))
	}
}
