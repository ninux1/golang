package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This is small project where we fetch the values from a database using 2 channels and 2
// goroutines.
// Database is simulated using go maps.

//Declare the mapping

var age = map[string]int{
	"foo": 40,
	"bar": 20,
	"abc": 56,
	"xyz": 43,
}

func main() {

	name1 := "xyz"
	name2 := "bar"

	c1 := make(chan int)
	c2 := make(chan int)

	go fetch(name1, c1)
	go fetch(name2, c2)

	fmt.Println("Age of", name1, <-c1) // This blocks untill there is a value on thi channel
	fmt.Println("Age of", name2, <-c2)

}

func fetch(name string, c chan int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	c <- age[name]
}
