package main

import (
	"fmt"
)

type emp struct {
	empno   int
	empname string
}

func main() {
	e := emp{203, "James"}
	printfunc(e)
}

func printfunc(r error) {
	fmt.Println(r.Error())
}

func (e emp) Error() string {
	return "Something went wrong"
}
