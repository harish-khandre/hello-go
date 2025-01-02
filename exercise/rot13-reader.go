package main

import (
	"io"
	"os"
	"strings"
)

// Define the rot13Reader type
type rot13Reader struct {
	r io.Reader
}

// Implement the Read method for rot13Reader
func (r *rot13Reader) Read(p []byte) (int, error) {
	// Read data from the underlying reader
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}

	// Apply ROT13 transformation to the read bytes
	for i := 0; i < n; i++ {
		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = 'A' + (p[i]-'A'+13)%26
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 'a' + (p[i]-'a'+13)%26
		}
	}

	return n, nil
}

func rot13ReaderExercise() {
	// Input string
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // "You cracked the code!" in ROT13

	// Wrap the input reader with rot13Reader
	r := rot13Reader{s}

	// Use io.Copy to decode the text and print it to stdout
	io.Copy(os.Stdout, &r)
}
