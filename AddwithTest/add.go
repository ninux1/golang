package main

import "fmt"

func main() {
	a := 10
	b := 10
	res := add(a, b)
	fmt.Printf("The sum is %d\n", res)
}

func add(x int, y int) int {
	return x + y
}
