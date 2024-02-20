package lfu_cache

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1))
	cache.Put(3, 3)
	t.Log(cache.Get(2))
	t.Log(cache.Get(3))
	cache.Put(4, 4)
	t.Log(cache.Get(1))
	t.Log(cache.Get(3))
	t.Log(cache.Get(4))
}
