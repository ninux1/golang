package main

import "fmt"

func main() {
	l := []int{12, 16, 18, 23, 45, 50, 63, 70, 87, 96, 112, 140}

	n := 63
	low := 0
	high := 11

	position := binsearch(l, n, low, high)

	fmt.Printf("The index of %d is %d\n", n, position)
}

func binsearch(l []int, n int, lowidx int, highidx int) int {
	mididx := (highidx + lowidx) / 2

	if n < l[mididx] {
		highidx = mididx - 1
		return binsearch(l, n, lowidx, highidx)

	}
	if n > l[mididx] {
		lowidx = mididx + 1
		return binsearch(l, n, lowidx, highidx)
	}
	return mididx
}
