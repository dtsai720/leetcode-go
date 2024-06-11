package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestContinuousSubarraySum(t *testing.T) {
	t.Parallel()

	t.Run("Test 1", func(t *testing.T) {
		t.Parallel()

		nums := []int{23, 2, 4, 6, 7}
		k := 6
		got := src.ContinuousSubarraySum(nums, k)
		assert.True(t, got)
	})

	t.Run("Test 2", func(t *testing.T) {
		t.Parallel()

		nums := []int{23, 2, 6, 4, 7}
		k := 6
		got := src.ContinuousSubarraySum(nums, k)
		assert.True(t, got)
	})

	t.Run("Test 3", func(t *testing.T) {
		t.Parallel()

		nums := []int{23, 2, 6, 4, 7}
		k := 13
		got := src.ContinuousSubarraySum(nums, k)
		assert.False(t, got)
	})
}
