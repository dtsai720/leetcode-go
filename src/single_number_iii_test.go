package src_test

import (
	"sort"
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestSingleNumberIII(t *testing.T) {
	t.Parallel()

	t.Run("TestSingleNumberIII", func(t *testing.T) {
		t.Parallel()

		nums := []int{1, 2, 1, 3, 2, 5}
		expected := []int{3, 5}
		result := src.SingleNumberIII(nums)
		sort.Ints(result)
		assert.Equal(t, expected, result)
	})

	t.Run("TestSingleNumberIII2", func(t *testing.T) {
		t.Parallel()

		nums := []int{0, -1}
		expected := []int{-1, 0}
		result := src.SingleNumberIII(nums)
		sort.Ints(result)
		assert.Equal(t, expected, result)
	})

	t.Run("TestSingleNumberIII3", func(t *testing.T) {
		t.Parallel()

		nums := []int{0, 1}
		expected := []int{0, 1}
		result := src.SingleNumberIII(nums)
		sort.Ints(result)
		assert.Equal(t, expected, result)
	})
}
