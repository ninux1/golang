package main

import (
	"fmt"
	"time"
)

func entrypoint() {

	word := "hello"

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(word)
	}()

	fmt.Println("Hello")
	word = "World"
	time.Sleep(3 * time.Second)
}

//func waitandprint(s string) {
//	fmt.Println(s)
//}
