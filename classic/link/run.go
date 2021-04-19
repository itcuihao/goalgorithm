package link

import "fmt"

type Link struct {
	Value int
	Next  *Link
}

// 收录链表相关题目
// https://www.jianshu.com/p/ce184b8c92a4
// k个有序链表merge成一个有序链表
// 二叉树转化为中序链表
// 链表排序

// 给定一个链表，其中奇数位是升序的，偶数位是降序的，实现链表的排序
// 将原始链表按奇偶分开（leetcode原题：奇偶链表） 2、偶链表是降序的，所以将其逆序（leetcode原题：翻转链表）3、合并两个有序链表（leetcode原题：合并两个有序链表） 时复杂度为O(n),空间复杂度为O（1）

func createLink(n int) *Link {
	l := &Link{}
	t := l
	for i := 1; i <= n; i++ {
		p := &Link{Value: i}
		t.Next = p
		t = p
	}
	return l.Next
}

func printLink(l *Link) {
	a := []int{}
	t := l
	for t != nil {
		a = append(a, t.Value)
		t = t.Next
	}
	fmt.Println(a)
}

func run() {
	l := createLink(5)
	printLink(l)
}
