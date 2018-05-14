package main

import "fmt"

func do() {

	n := 1 //Even if you are using defer the variables are calculated first than any statements

	// for the below statements, defer will pass on the control to the next statement to execute first.
	defer fmt.Println("First", n)
	n = 2
	defer fmt.Println("Second", n)
	n = 3
	defer fmt.Println("Third", n)
}

func main() {
	do()
}
