package t1

import (
	"testing"
)

func TestRun(t *testing.T) {
	a := []int{1, 1, 4, 6}
	b := []int{2, 3, 5, 7}
	t.Log(merge(a, b))
}
