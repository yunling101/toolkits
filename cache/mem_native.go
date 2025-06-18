package cache

import "sync"

type NativeCache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewNativeKVCache() *NativeCache {
	return &NativeCache{
		data: make(map[string]interface{}),
	}
}

func (c *NativeCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *NativeCache) Get(key string) (value interface{}, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok = c.data[key]
	return
}

func (c *NativeCache) GetAll() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data
}

func (c *NativeCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[key]; ok {
		delete(c.data, key)
	}
}
