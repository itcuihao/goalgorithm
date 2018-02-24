// Bubble Sort in Golang
// 算法复杂度 O(n)
// 在一趟遍历中，不断对相邻的元素进行比较排序，小的在前，大的在后，
// 大的元素会不断沉底。一趟遍历完成，最大的元素会被排在后方正确的位置上。
// 然后缩小排序范围，即去掉最后位置正确的元素，对剩余元素继续遍历，
// 直到范围不能再缩小为止，排序完成。
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99) - rand.Intn(99)
	}
	return slice
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
		fmt.Println(items)
	}
}
