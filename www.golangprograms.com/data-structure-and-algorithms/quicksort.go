// Quick Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(10)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)

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

func quicksort(a []int) []int {
	fmt.Println("000:", a)
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	// 基准元素
	pivot := rand.Int() % len(a)
	fmt.Println("pivot:", pivot)
	// 如果基准不再数组末尾，则把基准移到末尾
	if pivot != len(a) {
		a[pivot], a[right] = a[right], a[pivot]
	}
	fmt.Println("a:", a)

	// 然后从左到右（除了最后的基准元素），循环移动小于等于基准元素的元素到数组的开头，每次移动 left 自增 1，表示下一个小于基准元素将要移动到的位置。
	for i, _ := range a {
		if a[i] < a[right] {
			fmt.Println("left:i-", left, ":", i)
			a[left], a[i] = a[i], a[left]
			left++
			fmt.Println(i, " : ", a)
		}

	}

	fmt.Println("b:", a)
	// 循环结束后 left 所代表的的位置就是基准元素的所有摆放的位置。所以最后将基准元素所在位置（这里是 right）与 left 所代表的的位置的元素交换位置。
	a[left], a[right] = a[right], a[left]
	fmt.Println("c:", a)

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

// 分治法的基本思想是：将原问题分解为若干个规模更小但结构与原问题相似的子问题。递归地解这些子问题，然后将这些子问题的解组合为原问题的解。

// 利用分治法可将快速排序的分为三步：

// 在数据集之中，选择一个元素作为”基准”（pivot）。
// 所有小于”基准”的元素，都移到”基准”的左边；所有大于”基准”的元素，都移到”基准”的右边。这个操作称为分区 (partition) 操作，分区操作结束后，基准元素所处的位置就是最终排序后它的位置。
// 对”基准”左边和右边的两个子集，不断重复第一步和第二步，直到所有子集只剩下一个元素为止。
