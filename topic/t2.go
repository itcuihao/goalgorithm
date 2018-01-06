package main

import "fmt"

func main() {
	s := make([]int, 5) // 长度为5，容量为5
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
