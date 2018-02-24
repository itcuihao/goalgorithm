package main

import (
	"fmt"
)

// 一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：
// 1.从第一个元素开始，该元素可以认为已经被排序
// 2.取出下一个元素，在已经排序的元素序列中从后向前扫描
// 3.如果该元素（已排序）大于新元素，将该元素移到下一位置
// 4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
// 5.将新元素插入到该位置后
// 6.重复步骤2~5

// 平均来说插入排序算法复杂度为 O(n^{2})

func main() {
	us := []int{6, 1, 3, 9, 2, 7, 4, 8, 5}
	insertionsort(us)
	fmt.Println(us)
}

func insertionsort(s []int) {
	if len(s) < 2 {
		return
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
