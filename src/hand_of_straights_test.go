package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestHandOfStraights(t *testing.T) {
	t.Parallel()

	assert.True(t, src.HandOfStraights([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	assert.False(t, src.HandOfStraights([]int{1, 2, 3, 4, 5}, 4))
}
