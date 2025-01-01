package main

import (
	"fmt"
)

// Basic Interface Example
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type implementing Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Rectangle type implementing Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Function to demonstrate polymorphism
func PrintShapeDetails(s Shape) {
	fmt.Printf("Shape Details: Area = %.2f, Perimeter = %.2f\n", s.Area(), s.Perimeter())
}

// Composing Interfaces
type Logger interface {
	Log(message string)
}

type ErrorLogger interface {
	Logger
	LogError(error)
}

// ConsoleLogger implementing Logger
type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("LOG:", message)
}

func (cl ConsoleLogger) LogError(err error) {
	fmt.Println("ERROR:", err)
}

// Using Empty Interface
func PrintAnything(value interface{}) {
	fmt.Println("Value:", value)
}

// Mocking with Interfaces
type Database interface {
	Save(data string) error
}

type RealDatabase struct{}

func (db RealDatabase) Save(data string) error {
	fmt.Println("Saving data to the real database:", data)
	return nil
}

type MockDatabase struct {
	Storage []string
}

func (mdb *MockDatabase) Save(data string) error {
	mdb.Storage = append(mdb.Storage, data)
	fmt.Println("Mock Save:", data)
	return nil
}

func TestDatabase(db Database) {
	db.Save("Test Data")
}

func interfaceImplementation() {
	// Basic Interface Example
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 4}

	PrintShapeDetails(circle)
	PrintShapeDetails(rectangle)

	// Composing Interfaces
	logger := ConsoleLogger{}
	logger.Log("This is a log message.")
	logger.LogError(fmt.Errorf("%s", "This is the error logger"))

	// Using Empty Interface
	PrintAnything(42)
	PrintAnything("Hello, Go!")
	PrintAnything([]int{1, 2, 3})

	// Mocking with Interfaces
	realDB := RealDatabase{}
	mockDB := &MockDatabase{}

	fmt.Println("\nUsing Real Database:")
	TestDatabase(realDB)

	fmt.Println("\nUsing Mock Database:")
	TestDatabase(mockDB)
	fmt.Println("Mock Database Storage:", mockDB.Storage)
}

func typeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
