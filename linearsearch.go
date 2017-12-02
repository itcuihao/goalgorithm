// Linear Search in Golang
package main

import "fmt"

func linearsearch(datalist []int, element int) (b bool) {
	for _, item := range datalist {
		if item == element {
			b = true
			return
		}
	}
	return
}

func main() {
	items := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(linearsearch(items, 55))
}
