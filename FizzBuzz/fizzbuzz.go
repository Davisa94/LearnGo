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

func arrayPrinter(fizzyWord string) {
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

func main() {
	fmt.Println("e")
}
