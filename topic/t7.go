package main

import (
	"fmt"
	"time"
)

func main() {
	t := make(chan bool)
	go timeout(t)

	ch := make(chan string)
	go readword(ch)

	select {
	case w := <-ch:
		fmt.Println("Received", w)
	case <-t:
		fmt.Println("Timeout.")
	}
}

func readword(ch chan string) {
	fmt.Println("Type word & hit enter")
	var word string
	fmt.Scanf("%s", &word)
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- true
}
