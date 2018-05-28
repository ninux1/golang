package main

import (
	"fmt"
	"golang/counters"
)

func main() {
	c := counters.AlertCounter(10)
	fmt.Printf("Exported Counter initialized %d\n", c)
}
