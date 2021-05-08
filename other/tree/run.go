package tree

import "fmt"

type Tree struct {
	Val   int
	Right *Tree
	Left  *Tree
}

func run() {
	t := &Tree{
		Val: 1,
		Right: &Tree{
			Val:   2,
			Right: nil,
			Left:  nil,
		},
		Left: &Tree{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
	}
	a := rt(t)
	fmt.Println(a)
}

func rt(t *Tree) []int {
	if t == nil {
		return nil
	}
	a := make([]int, 0)
	tmp := make([]*Tree, 0)
	var tmpRoot *Tree
	for t != nil || len(tmp) > 0 {
		for t != nil {
			tmp = append(tmp, t.Left)
			t = t.Left
		}
		t = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		if t.Right == nil || t.Right == tmpRoot {
			a = append(a, t.Val)
			tmpRoot = t
			t = nil
		} else {
			tmp = append(tmp, t)
			t = t.Right
		}
	}

	return a
}

type link struct {
	Val  int
	Next *link
}

func linkq(l *link) bool {
	a := l
	b := l

	for b != nil {
		b = b.Next
		break
	}
	for b != nil {
		if a == b {
			return true
		}
		a = a.Next
		b = b.Next
	}
	return false
}
