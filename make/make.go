package main

import "fmt"

// This code is to demo make feature of Golang.
// make allocates memory in golang.
// Below the second argument to make is the actual len of slice.
// Below the third argument to make is the total capacity of the slice you may fill up slice upto.

func main() {
	//a := [5]int{} // Slice declaration
	a := make([]int, 5, 10)

	fmt.Println(len(a))

	for index, ele := range a {
		fmt.Println(index, ele)
	}

}
