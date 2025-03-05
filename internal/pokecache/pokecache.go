package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		mu:       sync.Mutex{},
	}

	go cache.readLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()         // Lock before modifying the map
	defer c.mu.Unlock() // Ensure unlock happens

	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock() // Lock before reading to ensure thread safety
	defer c.mu.Unlock()

	if entry, exists := c.entries[key]; exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()

		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.entries, key) // Remove expired entry
			}
		}
		c.mu.Unlock()
	}
}
