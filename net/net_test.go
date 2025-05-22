package net

import (
	"testing"
)

func TestGetStatusCode(t *testing.T) {
	r := NewRequest().GetStatusCode("http://www.baidu.com")
	if r.Error != nil {
		t.Error(r.Error)
	}
	t.Log(r.Data)
}
