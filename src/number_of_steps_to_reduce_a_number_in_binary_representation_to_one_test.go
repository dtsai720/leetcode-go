package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestNumberOfStepsToReduceANumberInBinaryRepresentationToOne(t *testing.T) {
	t.Parallel()

	t.Run("Test case 1", func(t *testing.T) {
		t.Parallel()

		s := "1101"
		expected := 6
		result := src.NumberOfStepsToReduceANumberInBinaryRepresentationToOne(s)
		assert.Equal(t, expected, result)
	})

	t.Run("Test case 2", func(t *testing.T) {
		t.Parallel()

		s := "10"
		expected := 1
		result := src.NumberOfStepsToReduceANumberInBinaryRepresentationToOne(s)
		assert.Equal(t, expected, result)
	})

	t.Run("Test case 3", func(t *testing.T) {
		t.Parallel()

		s := "1"
		expected := 0
		result := src.NumberOfStepsToReduceANumberInBinaryRepresentationToOne(s)
		assert.Equal(t, expected, result)
	})
}
