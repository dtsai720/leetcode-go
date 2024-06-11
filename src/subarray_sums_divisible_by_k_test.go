package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestSubarraySumsDivisibleByK(t *testing.T) {
	t.Parallel()

	t.Run("Test 1", func(t *testing.T) {
		t.Parallel()

		nums := []int{4, 5, 0, -2, -3, 1}
		k := 5
		expected := 7
		got := src.SubarraySumsDivisibleByK(nums, k)
		assert.Equal(t, expected, got)
	})

	t.Run("Test 2", func(t *testing.T) {
		t.Parallel()

		nums := []int{5}
		k := 9
		expected := 0
		got := src.SubarraySumsDivisibleByK(nums, k)
		assert.Equal(t, expected, got)
	})
}
