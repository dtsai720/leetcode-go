package src_test

import (
	"sort"
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestWordBreakII(t *testing.T) {
	t.Parallel()

	t.Run("Case I", func(t *testing.T) {
		t.Parallel()

		s := "catsanddog"
		wordDict := []string{"cat", "cats", "and", "sand", "dog"}
		expected := []string{"cat sand dog", "cats and dog"}
		output := src.WordBreakII(s, wordDict)

		sort.Strings(output)
		sort.Strings(expected)

		assert.Equal(t, expected, output)
	})

	t.Run("Case II", func(t *testing.T) {
		t.Parallel()

		s := "pineapplepenapple"
		wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
		expected := []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"}
		output := src.WordBreakII(s, wordDict)

		sort.Strings(output)
		sort.Strings(expected)

		assert.Equal(t, expected, output)
	})

	t.Run("Case III", func(t *testing.T) {
		t.Parallel()
		s := "catsandog"
		wordDict := []string{"cats", "dog", "sand", "and", "cat"}
		expected := []string{}
		output := src.WordBreakII(s, wordDict)

		sort.Strings(output)
		sort.Strings(expected)

		assert.Equal(t, expected, output)
	})
}
