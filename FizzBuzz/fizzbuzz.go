/**********************************************************
# Author: Austin Benitez
# Date: 2/2/22
# Assignment: Fizz Buzz
# Description:
# The rules: Count to 100*
# Every Multiple of 3 is replaced with Fizz
# Every multiple of 5 is replaced with Buzz
# Every multiple of both is replaced with FizzBuzz
# Stretch: every other output is indented
***********************************************************/
package main

import (
	"fmt"
	"strconv"
)

func arrayPrinter(fizzyWord []string) {
	for _, item := range fizzyWord {
		fmt.Println(item)
	}

}

func fizzBuzzLogic(num int) string {
	switch {
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(num)
	}

}

// Given a limit or size execute the logic on every number from 0
// to the size
func fizzyBuzz(size int) []string {
	fizzyArray := make([]string, size+1)
	for iteration := range fizzyArray {
		fizzyArray[iteration] = fizzBuzzLogic(iteration)
	}
	return fizzyArray
}

func main() {
	fizzyBuzzy := fizzyBuzz(100)
	fmt.Println("Not Formatted:\n")
	fmt.Println("========================\n")
	arrayPrinter(fizzyBuzzy)
	fmt.Println("========================\n")
}
