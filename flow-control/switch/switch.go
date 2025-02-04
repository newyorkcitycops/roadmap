package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s\n", os)
	}
}
