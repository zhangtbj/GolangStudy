package main

import "fmt"

func sender(ch chan bool) {
	fmt.Println("Hello World")
	ch <- true
}

func receiver(ch chan bool) bool {
	return <- ch
}

func main() {
	ch := make(chan bool)
	go func() {
		ch <- true
	}()
	fmt.Println(<-ch)
}