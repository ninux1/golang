package main

// This code is to demo the usage of garbage collection and memstats

// There are 2 memory we need to worry about
//			1) Stack which temporary to a function
//			2) Heap which is a global amount of memory that Go is using for objects which will have to be garbage collected
//			3) Stack dissappears once the function exits.
//			4) Anything that is not on the stack but is in a function would have escaped & left the function
//			   and will be on heap
//			5) use "go run -gcflags="-m" memstats.go

import (
	"fmt"
	"runtime"
	"time"
)

func length(s string) int {
	c := []byte(s)
	n := len(c)
	b := make([]byte, n)
	return len(b)
}

func main() {
	s := "Hello World!"
	s += s
	s += s
	s += s

	start := time.Now()
	for {
		if time.Since(start) > time.Second {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Printf("Heap size %d\n", r.HeapAlloc)
			fmt.Printf("NumGC %d\n", r.NumGC)
			start = time.Now()
		}
		length(s)
	}
}
