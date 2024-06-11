package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestSpecialArrayWithXElementsGreaterThanOrEqualX(t *testing.T) {
	t.Parallel()

	t.Run("Test1", func(t *testing.T) {
		t.Parallel()

		nums := []int{3, 5}
		expected := 2
		got := src.SpecialArrayWithXElementsGreaterThanOrEqualX(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("Test2", func(t *testing.T) {
		t.Parallel()

		nums := []int{0, 0}
		expected := -1
		got := src.SpecialArrayWithXElementsGreaterThanOrEqualX(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("Test3", func(t *testing.T) {
		t.Parallel()

		nums := []int{0, 4, 3, 0, 4}
		expected := 3
		got := src.SpecialArrayWithXElementsGreaterThanOrEqualX(nums)
		assert.Equal(t, expected, got)
	})
}
