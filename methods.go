// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// the Abs method has a receiver of type MethodInput named v.

package main

import (
	"fmt"
	"math"
)

type MethodInput struct {
	X, Y float64
}

type MethodPointersInput struct {
	X, Y float64
}

func (v MethodPointersInput) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *MethodPointersInput) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v MethodInput) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func method() {
	f := MyFloat(-math.Sqrt2)
	v := MethodInput{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(f.Abs())

	m := MethodPointersInput{3, 4}
	m.Scale(10)
	fmt.Println(m.Abs())
}
