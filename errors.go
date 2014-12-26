package flandmark

import "errors"

var (
	ErrCouldNotLoad = errors.New("Could not load Haar cascade.")
	ErrDataSize     = errors.New("Unexpected data size.")
)
