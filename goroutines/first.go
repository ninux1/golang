package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Hello How are you"

	// In golang threads are created by prefixing go in front of any code.
	go func() { // anonymous function calculates the function arguments first. since there are no function arguments no
		// arguments are calculated and modified msg string is used from the go routine.
		fmt.Println(msg)
	}()

	//go func(msg string) {
	//	fmt.Println(msg)
	//}(msg)

	msg = "Changed String now"

	//This delay is added to let the go routine finish execution and then exit the main context else we may not be able
	// to see what happened within the go routine.
	time.Sleep(100 * time.Millisecond)
}

//func printMsg(msg string) {
//	fmt.Println(msg)
//}
