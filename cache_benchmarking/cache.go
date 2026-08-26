package cachebenchmarking

import (
	"sync"
	"sync/atomic"
)

type Cache struct {
	mu         *sync.RWMutex
	data       map[string]interface{}
	hits       int64
	once       sync.Once
	defaultVal string
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]interface{}), mu: new(sync.RWMutex)}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get_With_Read_Lock(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	if ok {
		atomic.AddInt64(&c.hits, 1)
		return val
	}
	return ""

}

func (c *Cache) Get_With_Write_Lock(key string) any {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	if ok {
		c.hits += 1
		return val
	}

	return ""

}
