package time

import (
	"fmt"
	"testing"
)

func TestLastDateOfMonth(t *testing.T) {
	d := New().FirstDateOfWeek()
	fmt.Println(d)
}
