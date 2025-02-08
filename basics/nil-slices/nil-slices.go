package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("len=%d cap=%d %s\n", len(s), cap(s), s)
	if s == nil {
		fmt.Println("s is nil")
	}
}
