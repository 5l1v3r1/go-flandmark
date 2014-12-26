package flandmark

// #include "flandmark_binding.h"
// #cgo CXXFLAGS: -Iflandmark/libflandmark -I/usr/local/include/opencv -I/usr/include/opencv
// #cgo LDFLAGS: -lopencv_core -lopencv_imgproc -lopencv_objdetect flandmark/libflandmark/libflandmark_static.a
import "C"

import (
	"runtime"
	"unsafe"
)

// Image is an OpenCV-compatible image.
type Image struct {
	pointer unsafe.Pointer
}

// NewRGBAImage creates an image from raw RGBA data.
// The image will automatically be freed.
func NewRGBAImage(data []byte, width int, height int) *Image {
	if len(data) != 4 * width * height {
		return nil
	}
	buffer := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	v := C.flandmark_binding_image_rgba(buffer, C.int(width), C.int(height))
	res := &Image{v}
	runtime.SetFinalizer(res, res.free)
	return res
}

// NewGrayImage creates an image from raw grayscale data.
// The image will automatically be freed.
func NewGrayImage(data []byte, width int, height int) *Image {
	if len(data) != width * height {
		return nil
	}
	buffer := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	v := C.flandmark_binding_image_gray(buffer, C.int(width), C.int(height))
	res := &Image{v}
	runtime.SetFinalizer(res, res.free)
	return res
}

func (i *Image) free() {
	C.flandmark_binding_image_free(i.pointer)
}
