package t1

import (
	"fmt"
	"strconv"
	"strings"
)

func merge(a, b []int) []int {
	i, j := 0, 0
	la, lb := len(a), len(b)

	c := make([]int, 0, la+lb)

	for i < la && j < lb {

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

func maxNum() {
	nums := []int{3, 5, 9, 20, 22, 225}
	sort(nums)
	i := mergeN(nums)
	fmt.Println(i)
}

func sort(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if compare(nums[i], nums[j]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return
}

func compare(i, j int) bool {
	a := strconv.Itoa(i) + strconv.Itoa(j)
	b := strconv.Itoa(j) + strconv.Itoa(i)
	return a < b
}

func mergeN(nums []int) int {
	s := make([]string, 0, len(nums))
	for _, v := range nums {
		s = append(s, strconv.Itoa(v))
	}
	ss, err := strconv.Atoi(strings.Join(s, ""))
	if err != nil {
		return -1
	}
	return ss
}
