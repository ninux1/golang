package main

import "fmt"

func entrypoint2() {

	c := make(chan bool)
	go waitandprint(c)
	fmt.Println("Hello")
	c <- true // Entering true in the channels
}

func waitandprint(c chan bool) {

	if b := <-c; b { // waiting till we receive "true" on channel c
		fmt.Println("World")
	}

}
