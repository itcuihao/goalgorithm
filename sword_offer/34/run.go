package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	l := make([]int, 0, 4)
	find(root, target, l, &res)
	return res
}

func find(root *TreeNode, target int, l []int, res *[][]int) {
	if root == nil {
		return
	}
	l = append(l, root.Val)
	tmp := target - root.Val
	if root.Left == nil && root.Right == nil {
		if tmp == 0 {
			ll := make([]int, len(l))
			copy(ll, l)
			*res = append(*res, ll)
			fmt.Println("res; ", res)
		}
	}
	if root.Left != nil {
		find(root.Left, tmp, l, res)
	}
	if root.Right != nil {
		find(root.Right, tmp, l, res)
	}
	l = l[:len(l)-1]
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	l := pathSum(root, 22)
	fmt.Println(l)
}
