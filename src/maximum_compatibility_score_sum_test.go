package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestMaximumCompatibilityScoreSum(t *testing.T) {
	t.Parallel()

	t.Run("Example 1", func(t *testing.T) {
		t.Parallel()

		students := [][]int{
			{1, 1, 0},
			{1, 0, 1},
			{0, 0, 1},
		}
		mentors := [][]int{
			{1, 0, 0},
			{0, 0, 1},
			{1, 1, 0},
		}

		result := src.MaximumCompatibilityScoreSum(students, mentors)
		expected := 8
		assert.Equal(t, expected, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		t.Parallel()

		students := [][]int{
			{0, 0},
			{0, 0},
			{0, 0},
		}
		mentors := [][]int{
			{1, 1},
			{1, 1},
			{1, 1},
		}

		result := src.MaximumCompatibilityScoreSum(students, mentors)
		expected := 0
		assert.Equal(t, expected, result)
	})
}
