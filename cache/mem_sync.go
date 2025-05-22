package cache

import "sync"

type syncMap struct {
	data sync.Map
}

func NewKVCache() *syncMap {
	return &syncMap{
		data: sync.Map{},
	}
}

func (c *syncMap) Set(key, value any) {
	c.data.Store(key, value)
}

func (c *syncMap) Get(key any) (any, bool) {
	return c.data.Load(key)
}

func (c *syncMap) Delete(key any) {
	c.data.Delete(key)
}

func (c *syncMap) Range(f func(key, value any) bool) {
	c.data.Range(f)
}
