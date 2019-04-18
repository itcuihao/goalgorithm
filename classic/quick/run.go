package quick

// 三路快排
func quick3way(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	lt, i, gt := lo, lo+1, hi

	v := a[lo]

	for i <= gt {
		if a[i] < v {
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		} else if a[i] > v {
			a[gt], a[i] = a[i], a[gt]
			gt--
		} else {
			i++
		}
	}

	quick3way(a, lo, lt-1)
	quick3way(a, gt+1, hi)
}

// 普通快排
func quicksort(nums []int) {
	start := 0
	end := len(nums) - 1
	myquicksort(nums, start, end)
	return
}

func myquicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	p := partition(nums, start, end)
	myquicksort(nums, start, p-1)
	myquicksort(nums, p+1, end)
	return
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
