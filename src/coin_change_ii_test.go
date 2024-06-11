package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestCoinChangeII(t *testing.T) {
	t.Parallel()

	t.Run("TestCoinChangeII", func(t *testing.T) {
		t.Parallel()

		coins := []int{1, 2, 5}
		amount := 5
		expected := 4
		result := src.CoinChangeII(coins, amount)
		assert.Equal(t, expected, result)
	})

	t.Run("TestCoinChangeII", func(t *testing.T) {
		t.Parallel()

		coins := []int{2}
		amount := 3
		expected := 0
		result := src.CoinChangeII(coins, amount)
		assert.Equal(t, expected, result)
	})

	t.Run("TestCoinChangeII", func(t *testing.T) {
		t.Parallel()

		coins := []int{10}
		amount := 10
		expected := 1
		result := src.CoinChangeII(coins, amount)
		assert.Equal(t, expected, result)
	})
}
