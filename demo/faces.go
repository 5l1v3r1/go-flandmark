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
		fmt.Fprintln(os.Stderr, "Usage: censor <input_file>")
		os.Exit(1)
	}
	if err := errMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func errMain() error {
	r, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer r.Close()
	input, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	img, err := flandmark.GoGrayImage(input)
	if err != nil {
		return err
	}
	cascade, err := flandmark.LoadCascade("haarcascade_frontalface_alt.xml")
	if err != nil {
		return err
	}
	x, err := cascade.Detect(img, 1.1, 2, flandmark.Size{40, 40},
		flandmark.Size{1000000, 1000000})
	if err != nil {
		return err
	}
	fmt.Println("Got", len(x), "faces:", x)
	return nil
}
