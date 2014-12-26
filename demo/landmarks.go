package main

import (
	"fmt"
	"github.com/unixpickle/go-flandmark"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: landmarks <input_file>")
		os.Exit(1)
	}
	if err := errMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func errMain() error {
	// Read image
	img, err := readImage()
	if err != nil {
		return err
	}
	
	// Detect faces
	faces, err := detectFaces(img)
	if err != nil {
		return err
	}
	
	// Find facial features
	model, err := flandmark.LoadDefaultModel()
	if err != nil {
		return err
	}
	for _, face := range faces {
		features, err := model.Detect(img, face)
		if err != nil {
			return err
		}
		fmt.Println("Features for", face, "are", features)
	}
	
	return nil
}

func readImage() (*flandmark.Image, error) {
	r, err := os.Open(os.Args[1])
	if err != nil {
		return nil, err
	}
	defer r.Close()
	input, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return flandmark.GoGrayImage(input)
}

func detectFaces(img *flandmark.Image) ([]flandmark.Rect, error) {
	cascade, err := flandmark.LoadFaceCascade()
	if err != nil {
		return nil, err
	}
	return cascade.Detect(img, 1.1, 2, flandmark.Size{40, 40},
		flandmark.Size{1000000, 1000000})
}
