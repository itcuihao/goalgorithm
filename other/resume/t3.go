package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// var peo People = &Stduent{}
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// Go 方法集
// *T 接收 *t
// T  接收 *t/t
