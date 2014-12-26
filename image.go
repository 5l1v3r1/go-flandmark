package flandmark

import "image"

// GoRGBAImage uses the RGBA data from an image.Image to generate an Image which
// can be used throughout the library.
func GoRGBAImage(img image.Image) (*Image, error) {
	// Perform a rather inefficient conversion to RGBA data
	bounds := img.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y
	data := make([]byte, w*h*4)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x+bounds.Min.X, y+bounds.Min.Y).RGBA()
			i := x*4 + y*w*4
			data[i] = byte(r / 0x100)
			data[i+1] = byte(g / 0x100)
			data[i+2] = byte(b / 0x100)
			data[i+3] = byte(a / 0x100)
		}
	}
	return NewRGBAImage(data, w, h)
}

// GoGrayImage uses the grayscale data from an image.Image to generate an Image
// which can be used throughout the library.
func GoGrayImage(img image.Image) (*Image, error) {
	// Perform a rather inefficient conversion to grayscale data
	bounds := img.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y
	data := make([]byte, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := img.At(x+bounds.Min.X, y+bounds.Min.Y).RGBA()
			data[x+y*w] = byte((r + g + b) / 0x300)
		}
	}
	return NewGrayImage(data, w, h)
}
