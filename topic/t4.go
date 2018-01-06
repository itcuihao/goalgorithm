package main

import "fmt"

type people interface {
	Speak(string) string
}

type student struct{}

func (s *student) Speak(t string) (talk string) {
	if t == "b" {
		talk = "boy"
	} else {
		talk = "girl"
	}
	return
}

func main() {
	// var p people = student{}
	var p people = &student{}
	t := "b"
	fmt.Println(p.Speak(t))
}

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html
