package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestGetEqualSubstringsWithinBudget(t *testing.T) {
	t.Parallel()

	t.Run("Example 1", func(t *testing.T) {
		t.Parallel()

		maxCost := 3
		expected := 3
		result := src.GetEqualSubstringsWithinBudget("abcd", "bcdf", maxCost)
		assert.Equal(t, expected, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		t.Parallel()

		maxCost := 3
		expected := 1
		result := src.GetEqualSubstringsWithinBudget("abcd", "cdef", maxCost)
		assert.Equal(t, expected, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		t.Parallel()

		maxCost := 0
		expected := 1
		result := src.GetEqualSubstringsWithinBudget("abcd", "acde", maxCost)
		assert.Equal(t, expected, result)
	})
}
