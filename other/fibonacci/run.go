package fibonacci

func f(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return f(n-1) + f(n-2)
}

func fRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	h := 0
	hs := 1
	c := 0
	for i := 2; i <= n; i++ {
		c = h + hs
		h = hs
		hs = c
	}
	return c
}
