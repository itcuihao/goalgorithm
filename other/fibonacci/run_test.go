package fibonacci

import "testing"

func TestF(t *testing.T) {
	n := 10
	t.Log(f(n))
}
func TestFRecursive(t *testing.T) {
	n := 10
	t.Log(fRecursive(n))
}
