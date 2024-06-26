package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	t.Parallel()

	cache := src.NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1))
	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(2))
	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 4, cache.Get(4))
}
