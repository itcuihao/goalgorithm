package bigmul

import "strconv"

func mul(a, b string) string {
	if a == "" || b == "" {
		return "0"
	}
	if a == "0" || b == "0" {
		return "0"
	}
	la := len(a)
	lb := len(b)
	dig := make([]int, la+lb)
	for i := la - 1; i >= 0; i-- {
		for j := lb - 1; j >= 0; j-- {
			p := (a[i] - '0') * (b[j] - '0')
			pa := i + j
			pb := i + j + 1
			sum := int(p) + dig[pb]
			dig[pa] += sum / 10
			dig[pb] = sum % 10
		}
	}
	s := ""
	for _, v := range dig {
		if !(v == 0 && len(s) == 0) {
			s += strconv.Itoa(v)
		}
	}
	return s
}
