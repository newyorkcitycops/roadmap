package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("When's saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}
