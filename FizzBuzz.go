package main

import (
	"fmt"
)

func fizzBuzz(num int) {

	if num%5 == 0 && num%3 == 0 {
		fmt.Println("fizzbuzz")
	} else if num%5 == 0 {
		fmt.Println("buzz")
	} else if num%3 == 0  {
		fmt.Println("fizz")
	} else {
		fmt.Println("no dice")
	}
}

func main() {
	fizzBuzz(15)
}


