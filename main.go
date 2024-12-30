package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

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

type PercentageOfInputs struct {
	Value      float64
	TotalValue float64
}

func percentageOf(input PercentageOfInputs) string {
	if input.TotalValue == 0 {
		return "0%"
	}
	percentage := (input.Value / input.TotalValue) * 100
	return fmt.Sprintf("%.2f%%", percentage)
}

var c, python, java bool

func main() {
	// x := 7
	// y := 9
	//
	// fmt.Println(add(x, y))
	//
	// fmt.Println(mutiply(5, 2))
	//
	// fmt.Println(percentageOf(PercentageOfInputs{Value: 912.2, TotalValue: 10000}))
	//
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	//
	// var i int
	// fmt.Println(i, c, python, java)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	//
	// fmt.Printf("%v %v %v %q\n", x, y, a, b)
	//
	sum := 0

	for i := 0; i < 10; i++ {
		fmt.Println(sum)
		sum += i
	}

	fmt.Println(sum)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()

	// Switch without a condition is the same as switch true.

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

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
