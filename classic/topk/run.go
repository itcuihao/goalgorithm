package topk

// topk 问题
func topk(nums []int, k int) []int {
	if nums == nil {
		return nil
	}
	if len(nums) < k {
		return nums
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		i := partition(nums, start, end)
		switch {
		case i+1 < k:
			start = i + 1
		case i+1 > k:
			end = i - 1
		default:
			return nums[:k]
		}
	}
	return nil
}

func partition(nums []int, start, end int) int {
	povit := nums[start]
	s := start + 1
	e := end
	for s <= e {
		if nums[s] < povit && nums[e] > povit {
			nums[s], nums[e] = nums[e], nums[s]
			s++
			e--
		}
		if nums[s] >= povit {
			s++
		}
		if nums[e] <= povit {
			e--
		}
	}
	nums[start], nums[e] = nums[e], nums[start]
	return e
}
