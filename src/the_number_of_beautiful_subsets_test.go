package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestTheNumberOfBeautifulSubset(t *testing.T) {
	t.Parallel()

	t.Run("Test case 1", func(t *testing.T) {
		t.Parallel()

		nums := []int{2, 6, 4}
		k := 2
		expected := 4
		got := src.TheNumberOfBeautifulSubset(nums, k)
		assert.Equal(t, expected, got)
	})

	t.Run("Test case 2", func(t *testing.T) {
		t.Parallel()

		nums := []int{1}
		k := 1
		expected := 1
		got := src.TheNumberOfBeautifulSubset(nums, k)
		assert.Equal(t, expected, got)
	})
}
