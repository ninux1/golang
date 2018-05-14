package main

import "fmt"

func do() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
}

func main() {
	do()
}
