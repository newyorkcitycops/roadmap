package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	head := &List[int]{val: 1}

	second := &List[int]{val: 2}
	head.next = second

	third := &List[int]{val: 3}
	second.next = third

	current := head
	for current != nil {
		fmt.Println(current.val)
		current = current.next

	}
}
