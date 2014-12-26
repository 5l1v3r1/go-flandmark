package flandmark

import "errors"

var (
	ErrBadArgument  = errors.New("Invalid argument.")
	ErrCouldNotLoad = errors.New("Could not load file.")
	ErrDataSize     = errors.New("Got unexpected data size.")
	ErrDetect       = errors.New("Failed to detect.")
	ErrNormalize    = errors.New("Failed to normalize.")
)
