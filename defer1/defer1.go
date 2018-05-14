package main

import "fmt"

func do() {

	var counter int

	// Check this out. function in function (anonymous function call
	// The logic inside below anonymous function will be evaluated during exit of function do()
	// As there are no arguments for the below function , nothing is evaluated.
	defer func() {
		if counter != 0 {
			fmt.Println("From Anonymous function ", counter)
		}
	}()

	counter = 1
}

func main() {
	do()
}
