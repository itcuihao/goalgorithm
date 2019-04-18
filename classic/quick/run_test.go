package quick

import (
	"fmt"
	"testing"
)

func TestQuick3(t *testing.T) {
	a := []int{2, 1, 3, 5, 4}

	quick3way(a, 0, 4)
	fmt.Println(a)
}
func TestQuick2(t *testing.T) {
	a := []int{2, 1, 3, 5, 4}

	quicksort(a)
	t.Log(a)
}
func TestQuick1(t *testing.T) {
	a := []int{2, 1, 3, 5, 4}

	myquicksort(a, 0, 4)
	t.Log(a)
}
