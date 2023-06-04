package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val []byte
	createdAt time.Time
}

func newCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	entry, ok := c.cache[key]
	return entry.val, ok
}