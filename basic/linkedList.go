package main

import "fmt"

type list struct {
	name string
	next *list
}

func printList(list *list) {
	for list != nil {
		fmt.Println(list.name)
		list = list.next
	}
}

func insertTop(personList **list, person list) {
	person.next = *personList
	*personList = &person
}

func insertTail(list *list, insert list) {
	for list.next != nil {
		list = list.next
	}
	list.next = &insert
}

func main() {
	personList := &list{name: "jordan"}
	person2 := list{name: "ellen"}
	person3 := list{name: "jerry"}

	insertTop(&personList, person2)

	insertTail(personList, person3)

	printList(personList)
}