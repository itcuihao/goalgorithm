package bigmul

import "testing"

func TestMul(t *testing.T) {
	a := "2222222222222222222222222222222222"
	b := "333333333333333333333333333333333333"
	c := mul(a, b)
	t.Log(c)
}
