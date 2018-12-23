package quick

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
