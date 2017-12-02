// 选择排序
// 工作原理：假定第一个元素是最小（大）的，然后从剩余的元素中寻找最小（大）的，
// 然后放在已排序的序列末尾，反复，直到序列结束。
// 比较次数O(n2),当n值较小的时候选择排序比冒泡排序快
// 优点：移动元素原地操作
package main

import "fmt"

func main() {
	recresivesort := []int{2, 4, 3, 5, 6, 1, 7, 9, 8}
	recursiveSelectionSort(recresivesort, 0, len(recresivesort))
	fmt.Println(recresivesort)
	sort := []int{2, 4, 3, 5, 6, 1, 7, 9, 8}
	selectionSorts(sort)
	fmt.Println(sort)
}

// 递归实现选择排序
func recursiveSelectionSort(s []int, start, end int) {
	if start+1 == end {
		return
	}
	min := start
	for i := start + 1; i < end; i++ {
		if s[min] > s[i] {
			min = i
		}
	}
	if min != start {
		// 数值间利用两个变量本身进行交换
		s[min], s[start] = s[start], s[min]
	}

	recursiveSelectionSort(s, start+1, end)
}

func selectionSort(s []int) {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i; j < len(s); j++ {
			if s[j] < s[min] {
				s[j], s[min] = s[min], s[j]
			}
		}
	}
}

func selectionSorts(s []int) {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
}
