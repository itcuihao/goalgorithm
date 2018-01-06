package main

import "fmt"

// defer是在函数末尾的return前执行，先进后执行，具体见问题1。
// 函数调用时 int 参数发生值拷贝。

func calc(i string, a, b int) int {
	r := a + b
	fmt.Println(i, a, b, r)
	return r
}

func main() {
	a := 1
	b := 2
	// defer 只对最后一个函数有效
	defer calc("1", a, calc("10", a, b))

	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
