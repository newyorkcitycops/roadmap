package main

import "fmt"

func square(x *int) {
	*x = *x * 2
}

func main() {
	x := 2
	square(&x)
	fmt.Println(x)
}
