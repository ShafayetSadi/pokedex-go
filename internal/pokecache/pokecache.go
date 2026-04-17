package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	val        []byte
	created_at time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		store: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.store[key] = cacheEntry{
		val:        val,
		created_at: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	value, ok := c.store[key]
	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, entry := range c.store {
		if entry.created_at.Before(now.Add(-last)) {
			delete(c.store, key)
		}
	}
}
