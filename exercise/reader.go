package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A' // Fill the slice with ASCII 'A'
	}
	return len(b), nil // Return the length of the slice and no error
}

func readerExercise() {
	reader.Validate(MyReader{})
}
