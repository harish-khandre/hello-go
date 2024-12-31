package main

import "fmt"

func slice() {
	e := []int{2, 3, 5, 7, 11, 13}
	printSlice(e)

	// Slice the slice to give it zero length.
	e = e[:0]
	printSlice(e)

	// Extend its length.
	e = e[:4]
	printSlice(e)

	// Drop its first two values.
	e = e[2:]
	printSlice(e)

	y := make([]int, 0, 5)
	printSlices("b", y)
}

func printSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
