package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image defines a custom image type with width, height, and a base color.
type Image struct {
	width, height int   // Dimensions of the image
	color         uint8 // Base color value for generating pixel colors
}

// Bounds defines the rectangular boundary of the image.
func (i Image) Bounds() image.Rectangle {
	// Returns a rectangle representing the dimensions of the image.
	return image.Rect(0, 0, i.width, i.height)
}

// ColorModel specifies the color model used by the image.
func (i Image) ColorModel() color.Model {
	// Returns the RGB color model used to interpret colors.
	return color.RGBAModel
}

// At determines the color of the pixel at (x, y).
func (i Image) At(x, y int) color.Color {
	// Returns a color.RGBA value based on the position (x, y).
	// The red and green components are adjusted by x and y, respectively.
	// Blue is fixed at 255, and alpha is fully opaque (255).
	return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}

func imagesExercise() {
	m := Image{100, 100, 100}
	pic.ShowImage(m)
}
