package tt

import "fmt"

// 无序数组的唯一元素
// 相关题目 leetcode 136 540

// 形如这样的无序数组 [4,4,2,2,3,1,1,5,5,9,9,6,6] ，
// 除了一个数字只会出现一次，其他数字成对出现而且相邻排列，
// 找出只出现一次的这个数字。
// 要求 时间复杂度 和 空间复杂度 尽可能低
// 测试样例：
// 输入: [1]，输出: 1
// 输入: [1,2,2]，输出: 1
// 输入: [1,1,2]，输出: 2
// 输入: [4,4,2,3,3,1,1,5,5,9,9,6,6]，输出: 2
// 输入: [4,4,2,2,1,1,5,5,9,9,3,6,6]，输出: 3
// 输入: [1,1,8,8,4,4,5,5,6,6,7,7,3,3,2]，输出: 2

func run() {
	nums := []int{4, 4, 2, 2, 3, 1, 1, 5, 5, 9, 9, 6, 6}
	i := search(nums)
	fmt.Println(i)
}
func search(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	s, e := 0, len(nums)
	for s < e {
		k := s + (e-s)/2
		fmt.Println("sek", s, e, k)
		if nums[k] != nums[k-1] && nums[k] != nums[k+1] {
			return nums[k]
		}
		fmt.Println(nums[:k])
		if nums[k] == nums[k-1] {
			if len(nums[:k])%2 == 0 {
				e = k
			}
		}
		if nums[k] == nums[k+1] {
			s = k
		}
	}
	return nums[e]
}

func sum(nums []int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
