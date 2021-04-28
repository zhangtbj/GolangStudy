package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2 * time.Second)
}

func sayHelloWorld() {
	fmt.Println("Hello World")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
	go sayHelloWorld()
	time.Sleep(1 * time.Second)
}
