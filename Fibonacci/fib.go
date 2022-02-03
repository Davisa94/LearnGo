package main

import (
	"fmt"
)

func fibonacci(num int) int {
	// If the number is less than zero invalid input
	fmt.Println(num)
	if num < 0 {
		fmt.Println("invalid number")
		return -1
	} else if num == 0 {
		return 0
	} else if num == 1 || num == 2 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func main() {
	fmt.Println(fibonacci(10))
}
