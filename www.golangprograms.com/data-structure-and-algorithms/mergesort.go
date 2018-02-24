package main

// 归并排序（英语：Merge sort，或mergesort），是创建在归并操作上的一种有效的排序算法，效率为  O(n\log n)。1945年由约翰·冯·诺伊曼首次提出。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用，且各层分治递归可以同时进行

// 迭代法
// 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
// 设定两个指针，最初位置分别为两个已经排序序列的起始位置
// 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
// 重复步骤3直到某一指针到达序列尾
// 将另一序列剩下的所有元素直接复制到合并序列尾

// 递归法
// 原理如下（假设序列共有 n个元素）：
// 将序列每相邻两个数字进行归并操作，形成  ceil(n/2)个序列，排序后每个序列包含两/一个元素
// 若此时序列数不是1个则将上述序列再次归并，形成 ceil(n/4)个序列，每个序列包含四/三个元素
// 重复步骤2，直到所有元素排序完毕，即序列数为1

// 最坏时间复杂度	(n\log n)
// 最优时间复杂度	(n)
// 平均时间复杂度	(n\log n)
// 空间复杂度	  (n)

import "fmt"

func main() {
	us := []int{6, 1, 3, 9, 2, 7, 4, 8, 5}
	s := mergeSort(us)
	fmt.Println(s)
}

func mergeSort(s []int) []int {
	fmt.Println("s:", s)
	var num = len(s)
	if num == 1 {
		return s
	}
	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)

	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = s[i]
		} else {
			right[i-middle] = s[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}
func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
