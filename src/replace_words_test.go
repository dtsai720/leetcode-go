package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestReplaceWords(t *testing.T) {
	t.Parallel()

	t.Run("Test 1", func(t *testing.T) {
		t.Parallel()

		dictionary := []string{"cat", "bat", "rat"}
		sentence := "the cattle was rattled by the battery"
		want := "the cat was rat by the bat"
		got := src.ReplaceWords(dictionary, sentence)
		assert.Equal(t, want, got)
	})

	//nolint:dupword
	t.Run("Test 2", func(t *testing.T) {
		t.Parallel()

		dictionary := []string{"a", "b", "c"}
		sentence := "aadsfasf absbs bbab cadsfafs"
		want := "a a b c"
		got := src.ReplaceWords(dictionary, sentence)
		assert.Equal(t, want, got)
	})
}
