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

func rt(root *Tree) []int {
	var a []int
	stk := []*Tree{}
	var prev *Tree
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			a = append(a, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
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

	for b.Next != nil && b.Next.Next != nil {
		if a == b {
			return true
		}
		a = a.Next
		b = b.Next.Next
	}
	return false
}
