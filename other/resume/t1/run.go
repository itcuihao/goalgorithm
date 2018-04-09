package t1

import (
	"fmt"
)

func merge(a, b []int) []int {
	i, j := 0, 0
	la, lb := len(a), len(b)

	c := make([]int, 0, la+lb)

	for i != la && j != lb {

		fmt.Println("i:j--", i, j)
		if a[i] > b[j] {
			c = append(c, b[j])
			fmt.Println("b:", c)
			j++
		} else {
			c = append(c, a[i])
			fmt.Println("a:", c)
			i++
		}
	}

	if i != la {
		c = append(c, a[i:]...)
	}

	if j != lb {
		c = append(c, b[j:]...)
	}
	return c
}
