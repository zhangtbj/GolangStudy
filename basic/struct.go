package main

import "fmt"

type person struct {
	name string
	age int
	title string
}

func main() {
	jordan := person {
		name: "jordan",
		age: 36,
		title: "engineer",
	}

	// 返回的是指向结构体的指针
	ellen := new(person)
	ellen.name = "ellen"
	ellen.age = 37
	ellen.title = "teacher"

	fmt.Println(jordan)
	fmt.Println(ellen)
}
