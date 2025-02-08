package main

import "fmt"

type Car struct {
	Driver   string
	Velocity float64
	Fuel     float64
}

func main() {
	car := Car{
		Driver:   "Tobias",
		Velocity: 42,
		Fuel:     50,
	}

	fmt.Println(car.Driver)
}
