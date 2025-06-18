package cache

import "sync"

type SyncMap struct {
	data sync.Map
}

func NewKVCache() *SyncMap {
	return &SyncMap{
		data: sync.Map{},
	}
}

func (c *SyncMap) Set(key, value any) {
	c.data.Store(key, value)
}

func (c *SyncMap) Get(key any) (any, bool) {
	return c.data.Load(key)
}

func (c *SyncMap) Delete(key any) {
	c.data.Delete(key)
}

func (c *SyncMap) Range(f func(key, value any) bool) {
	c.data.Range(f)
}
