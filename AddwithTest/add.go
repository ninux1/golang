package main

import "fmt"

func main() {
	a := 10
	b := 10
	res := Add(a, b)
	fmt.Printf("The sum is %d\n", res)
}

// Add function
func Add(x int, y int) int {
	return x + y
}
