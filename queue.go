package main

import "fmt"

type node struct {
	value int
}

func (n node) String() string {
	return fmt.Sprint(n.value)
}

type queue struct {
	nodes []*node
	size  int
	count int
	head  int
	tail  int
}

func (q *queue) Push(n *node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*node, 0, len(q.nodes)+q.size)
		// copy(nodes, q.nodes[q.head:])
		// copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		nodes = append(nodes, q.nodes...)
		nodes = append(nodes, n)
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
	fmt.Println(fmt.Sprintf("push:%+v", q))
}

func (q *queue) Pop() *node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	fmt.Println(fmt.Sprintf("pop:%+v", q))
	return node
}

func main() {
	closureQueue := func(size int) *queue {
		return &queue{
			nodes: make([]*node, size),
			size:  size,
		}
	}

	cq := closureQueue(1)
	cq.Push(&node{1})
	cq.Push(&node{2})
	fmt.Println(cq.Pop())
	cq.Push(&node{3})
	fmt.Println(cq.Pop())
	fmt.Println(cq.Pop())
	fmt.Println(cq.Pop())
}
