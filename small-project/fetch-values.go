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

	rand.Seed(time.Now().UnixNano())

	name1 := "xyz"
	name2 := "bar"

	c1 := make(chan int)
	c2 := make(chan int)

	go fetch(name1, c1)
	go fetch(name2, c2)

	select {
	case a := <-c1:
		fmt.Println("Age of", name1, a)
	case a := <-c2:
		fmt.Println("Age of", name2, a)
	case <-time.After(6 * time.Second):
		fmt.Println("Timed Out")
	}
}

func fetch(name string, c chan int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Simulating some delay
	c <- age[name]
}
