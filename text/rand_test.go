package text

import (
	"github.com/yunling101/toolkits/core"
	"testing"
)

func TestGenerateRandomToInt(t *testing.T) {
	s := make([]int, 0)
	n := make([]int, 0)
	for i := 0; i < 1000; i++ {
		v := GenerateRandomToInt(8)
		if core.IntContains(s, v) {
			n = append(n, v)
		}
		s = append(s, v)
	}
	if len(n) > 0 {
		t.Error("重复", n)
	}
}
