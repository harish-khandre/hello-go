package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java bool

func main() {
	x := 7
	y := 9

	fmt.Println(add(x, y))

	fmt.Println(mutiply(5, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var i int
	fmt.Println(i, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("%v %v %v %q\n", x, y, a, b)

	// A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world")
	fmt.Println("hello")

	// The * operator denotes the pointer's underlying value.
	i, j := 42, 2701

	p := &i // point to i

	fmt.Println(*p) // read i through the pointer

	// The & operator generates a pointer to its operand.
	*p = 21 // set i through the pointer

	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func add(x int, y int) int {
	return x + y
}

func substract(x int, y int) int {
	return x - y
}

func mutiply(x, y int) string {
	return fmt.Sprintf("%d", x*y)
}

func swap(x, y string) (string, string) {
	return y, x
}
