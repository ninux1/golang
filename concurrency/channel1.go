package main

import "fmt"

func entry1() {
	c := make(chan string, 5)

	//go fillinhello(c, 5)

	for i := 0; i < 5; i++ {
		c <- "Hello"
	}
	close(c)

	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-c)
	//}

	for ele := range c {
		fmt.Println(ele)
	}
}

/*
func fillinhello(c chan string, iter int) {
	for i := 0; i < 10; i++ {
		c <- "Hello"
	}
	close(c)
}
*/
