package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func methods() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.X = 1e9

	fmt.Println(v.X, v.Y)
	fmt.Println(v)
	fmt.Println(v1, p, v2, v3)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	s := primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	x := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(x, b)
	fmt.Println(names)

	//   This is an array literal:
	//
	// [3]bool{true, true, false}

	// And this creates the same array as above, then builds a slice that references it:
	// []bool{true, true, false}

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	w := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(w)

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

	var t []int
	fmt.Println(t, len(t), cap(t))

	if t == nil {
		fmt.Println("nil!")
	}

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

func ticTacToeBoard() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pew := make([]int, 10)
	for i := range pew {
		pew[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pew {
		fmt.Printf("%d\n", value)
	}
}

type Vertexx struct {
	Lat, Long float64
}

// var m map[string]Vertexx

func mapp() {
	// m = make(map[string]Vertexx)
	// m := map[string]Vertexx{
	// 	"Bell Labs": {
	// 		40.68433, -74.39967,
	// 	},
	// 	"Google": {
	// 		37.42202, -122.08408,
	// 	},
	// }

	m := map[string]Vertexx{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m["Bell Labs"])
}
