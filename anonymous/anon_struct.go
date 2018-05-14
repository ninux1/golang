// This code is to demo the Anonymous struct concept of golang.

package main

import (
	"fmt"
)

// This is a usual struct daclaration
//type construct struct {
//	shape string
//	area  float64
//}

func main() {
	//a := construct{"circle", 5.67} is a usual way of initializing the struct declared with name above

	//Below is an anonymous struct and initialized there itself.
	a := struct {
		shape string
		area  float64
	}{"circle", 8.33}

	//If you declare the anonymous struct as below then there is an explicit initialization as below.
	var b struct {
		vehicle string
		build   string
	}

	b.vehicle = "Toyota"
	b.build = "Land"

	var counter int // If there is no value assigned it will take 0 by default
	counter = 20

	pet := "Lion"

	fmt.Println("Anonymous struct :", a)
	fmt.Println("Anonymous struct :", b)
	fmt.Println("Counter Value :", counter)
	fmt.Println("My Pet is :", pet)
}
