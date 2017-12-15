package main

import "fmt"

// node 节点
type node struct {
	value int
}

func (n *node) String() string {
	return fmt.Sprint(n.value)
}

// stack 栈
type stack struct {
	nodes []*node // 元素列表
	count int     // 元素数量
}

func (s *stack) Push(n *node) {
	nodes := make([]*node, 0, s.count+1)
	nodes = append(s.nodes[:s.count], n)

	s.nodes = nodes
	s.count++
}

func (s *stack) Pop() *node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// 匿名函数
func anonymousNewStack() func() *stack {
	return func() *stack {
		return &stack{}
	}
}

func main() {

	// 闭包函数
	var closureNewStack = func() *stack {
		return &stack{}
	}
	cns := closureNewStack()
	cns.Push(&node{1})
	cns.Push(&node{2})
	fmt.Println(cns.Pop())
	cns.Push(&node{3})
	fmt.Println(cns.Pop())
	fmt.Println(cns.Pop())
	fmt.Println(cns.Pop())
	fmt.Println(cns.Pop())

	// 匿名函数
	ans := anonymousNewStack()()
	ans.Push(&node{1})
	ans.Push(&node{2})
	fmt.Println(ans.Pop())
	ans.Push(&node{3})
	fmt.Println(ans.Pop())
	fmt.Println(ans.Pop())
	fmt.Println(ans.Pop())
	fmt.Println(ans.Pop())
}
