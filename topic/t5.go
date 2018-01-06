package main

import "fmt"
import "reflect"

type ps interface {
	show()
}

type ss struct{}

func (s *ss) show() {

}

func live() ps {

	var s *ss
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
	return s
}

func main() {
	if live() == nil {
		fmt.Println("AAAA")
	} else {
		fmt.Println("BBBB")
	}
}
