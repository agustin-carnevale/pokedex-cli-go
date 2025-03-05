package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)

	if cache.entries == nil {
		t.Error("cache entries not initialized")
	}
}

func TestCreateCacheEntries(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 is not defined")
	}
	if string(actual) != "val1" {
		t.Error("key1 value doesn't match. Expected: val1")
	}
}
