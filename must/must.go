package main

import (
	"fmt"
	"regexp"
)

// This code is to demo must construct in golang.
// MustCompile is mostly used for safe initialization of global variables holding compiled regular exp.
var number = regexp.MustCompile("[0-9]+")

func main() {
	//re, err := regexp.Compile("[0-9+")
	//if err != nil {
	//	log.Fatalf("Bad Regular Expression Given %s", err)
	//}
	//fmt.Println(re.MatchString("ninux1078"))
	fmt.Println(number.MatchString("ninux1078"))
}
