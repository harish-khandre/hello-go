package main

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := range p {
		p[y] = make([]uint8, dx)
		for x := range p[y] {
			p[y][x] = uint8((x + y) / 2)
		}
	}
	return p
}

func _Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}
	for y := range p {
		for x := range p[y] {
			p[y][x] = uint8(x ^ y)
			// p[y][x] = uint8(x*y)
			// p[y][x] = uint8((x+y)/2)
		}
	}
	return p
}
