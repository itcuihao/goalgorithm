package main

import "fmt"

func main() {
	us := []int{6, 1, 3, 9, 2, 7, 4, 8, 5}
	pancakesort(us)
	fmt.Println(us)
}

func pancakesort(s []int) {
	for uns := len(s) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, s[0]
		for i := 1; i <= uns; i++ {
			if s[i] > lg {
				lx, lg = i, s[i]
			}
		}
		// move to final position in two flips
		flip(lx, s)
		flip(uns, s)
	}
}

func flip(r int, s []int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
