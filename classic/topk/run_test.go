package topk

import "testing"

func TestTopk(t *testing.T) {
	nums := []int{1, 2, 34, 2, 1, 5, 6, 2, 3, 56, 9}
	k := 4

	n := topk(nums, k)
	t.Log(n)
}
