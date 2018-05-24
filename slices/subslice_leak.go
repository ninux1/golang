package main

import "fmt"

func main() {

	subslice := testsubslice()
	fmt.Println(subslice)
	fmt.Println(cap(subslice))
	fmt.Println(subslice[:cap(subslice)])
}

func testsubslice() []int {
	s := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	return s[1:4]
}
