package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	printMap(m)

	m["Answer"] = 48
	printMap(m)

	delete(m, "Answer")

	v, ok := m["Answer"]
	fmt.Printf("The value is: %d. Is it present: %v", v, ok)
}

func printMap(m map[string]int) {
	fmt.Println("The value is:", m["Answer"])
}
