package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender rune
}

func main() {
	person := Person{
		Name:   "Person",
		Age:    69,
		Gender: 'M',
	}

	fmt.Println(person)
}
