package main

import "fmt"

func do() {

	n := 1 //Even if you are using defer, the arguments are calculated immediately

	// for the below statements, defer will execute at the exit of the function do()
	defer fmt.Println("First", n)
	n = 2
	defer fmt.Println("Second", n)
	n = 3
	defer fmt.Println("Third", n)
}

func main() {
	do()
}
