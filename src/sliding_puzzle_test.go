package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestSlidingPuzzle(t *testing.T) {
	t.Parallel()

	t.Run("case I", func(t *testing.T) {
		t.Parallel()

		board := [][]int{
			{1, 2, 3},
			{4, 0, 5},
		}
		expected := 1
		result := src.SlidingPuzzle(board)
		assert.Equal(t, expected, result)
	})

	t.Run("case II", func(t *testing.T) {
		t.Parallel()

		board := [][]int{
			{1, 2, 3},
			{5, 4, 0},
		}
		expected := -1
		result := src.SlidingPuzzle(board)
		assert.Equal(t, expected, result)
	})

	t.Run("case III", func(t *testing.T) {
		t.Parallel()

		board := [][]int{
			{4, 1, 2},
			{5, 0, 3},
		}
		expected := 5
		result := src.SlidingPuzzle(board)
		assert.Equal(t, expected, result)
	})
}
