package redis

import (
	"testing"
)

func TestCache(t *testing.T) {
	c := NewRedis("localhost:6379", "", 0)
	defer func() {
		_ = c.Close()
	}()

	err := c.Set("t1", "t1")
	if err != nil {
		t.Error(err)
		return
	}

	value, err := c.Get("t1")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(value)
}
