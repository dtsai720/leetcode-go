package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestMostProfitAssigningWork(t *testing.T) {
	t.Parallel()

	t.Run("Test case 1", func(t *testing.T) {
		t.Parallel()
		difficulty := []int{2, 4, 6, 8, 10}
		profit := []int{10, 20, 30, 40, 50}
		worker := []int{4, 5, 6, 7}
		expected := 100
		got := src.MostProfitAssigningWork(difficulty, profit, worker)
		assert.Equal(t, expected, got)
	})

	t.Run("Test case 2", func(t *testing.T) {
		t.Parallel()
		difficulty := []int{85, 47, 57}
		profit := []int{24, 66, 99}
		worker := []int{40, 25, 25}
		expected := 0
		got := src.MostProfitAssigningWork(difficulty, profit, worker)
		assert.Equal(t, expected, got)
	})
}
