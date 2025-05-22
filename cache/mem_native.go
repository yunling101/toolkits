package cache

import "sync"

type native struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewNativeKVCache() *native {
	return &native{
		data: make(map[string]interface{}),
	}
}

func (c *native) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *native) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *native) GetAll() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}

func (c *native) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
