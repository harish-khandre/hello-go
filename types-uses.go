package main

import (
	"fmt"
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

	var t []int
	fmt.Println(t, len(t), cap(t))

	if t == nil {
		fmt.Println("nil!")
	}
}

// var m map[string]Vertexx
