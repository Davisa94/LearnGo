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

func arrayPrinter(fizzyWords []string) {
	for _, item := range fizzyWords {
		fmt.Println(item)
	}

}

func staggeredArrayPrinter(fizzyWords []string) {
	for i := range fizzyWords {
		if i%2 == 0 {
			fmt.Println("\t", fizzyWords[i])
		} else {
			fmt.Println(fizzyWords[i])
		}
	}
}

func fizzBuzzLogic(num int) string {
	var temp = ""
	if num%3 == 0 {
		temp += "Fizz"
	}
	if num%5 == 0 {
		temp += "Buzz"
	}
	if temp == "" {
		return strconv.Itoa(num)
	}
	return temp

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
	fmt.Println("Formatted:\n")
	fmt.Println("========================\n")
	staggeredArrayPrinter(fizzyBuzzy)
	fmt.Println("========================\n")

}
