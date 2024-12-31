package main

import "fmt"

type Vertexx struct {
	Lat, Long float64
}

func object() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// If key is in m, ok is true. If not, ok is false.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// m = make(map[string]Vertexx)
	//
	// m := map[string]Vertexx{
	// 	"Bell Labs": {
	// 		40.68433, -74.39967,
	// 	},
	// 	"Google": {
	// 		37.42202, -122.08408,
	// 	},
	// }

	n := map[string]Vertexx{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(n)
}
