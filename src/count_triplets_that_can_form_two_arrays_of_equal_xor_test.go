package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestCountTripletsThatCanFormTwoArraysOfEqualXor(t *testing.T) {
	t.Parallel()

	t.Run("Test Case 1", func(t *testing.T) {
		t.Parallel()

		nums := []int{2, 3, 1, 6, 7}
		expected := 4
		got := src.CountTripletsThatCanFormTwoArraysOfEqualXor(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		t.Parallel()

		nums := []int{1, 1, 1, 1, 1}
		expected := 10
		got := src.CountTripletsThatCanFormTwoArraysOfEqualXor(nums)
		assert.Equal(t, expected, got)
	})
}
