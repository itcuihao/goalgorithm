package main

import "fmt"

// 在泡沫排序中，只比较阵列中相邻的二项，即比较的二项的间距（Gap）是1，梳排序提出此间距其实可大于1，改自插入排序的希尔排序同样提出相同观点。梳排序中，开始时的间距设定为阵列长度，并在循环中以固定比率递减，通常递减率设定为1.3。在一次循环中，梳排序如同泡沫排序一样把阵列从首到尾扫描一次，比较及交换两项，不同的是两项的间距不固定于1。如果间距递减至1，梳排序假定输入阵列大致排序好，并以泡沫排序作最后检查及修正。
// 递减率的设定影响着梳排序的效率，原作者以随机数作实验，得到最有效递减率为1.3的。如果此比率太小，则导致一循环中有过多的比较，如果比率太大，则未能有效消除阵列中的乌龟。

func main() {
	us := []int{6, 1, 3, 9, 2, 7, 4, 8, 5}
	combsort(us)
	fmt.Println(us)
}

func combsort(s []int) {
	var (
		n       = len(s)
		gap     = len(s)
		shrink  = 1.3
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if s[i] > s[i+gap] {
				s[i+gap], s[i] = s[i], s[i+gap]
				fmt.Println(s)
				swapped = true
			}
		}
	}
}
