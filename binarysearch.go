// Binary Search in Golang
// 二分搜索也称为折半搜索，从有序数组中查找指定元素。时间复杂度 O(log n)，空间复杂度 O(1)。
// 步骤：
// 给予一个包含n个带值元素的A数组以及目标值T
// 令L为0，R为n-1
// 如果L>R,则搜索以失败告终
// 令m（中间值元素）为 (L+R)/2
// 如果 Am < T，令L 为m+1 并返回步骤二
// 如果 Am > T, 令R 为m-1 并返回步骤二
// 当 Am = T，搜索结束，回传值 m

package main

import "fmt"

func binarySearch(element int, stack []int) (b bool) {

	low := 0
	high := len(stack) - 1

	for low <= high {
		median := (low + high) / 2
		if stack[median] < element {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(stack) || stack[low] != element {
		return
	}
	b = true
	return
}

func main() {
	items := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(binarySearch(76, items))
}
