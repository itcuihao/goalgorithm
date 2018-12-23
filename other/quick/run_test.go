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
