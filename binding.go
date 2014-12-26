package flandmark

// #include "flandmark_binding.h"
// #cgo CXXFLAGS: -Iflandmark/libflandmark -I/usr/local/include/opencv -I/usr/include/opencv
// #cgo LDFLAGS: -lopencv_core -lopencv_imgproc -lopencv_objdetect flandmark/libflandmark/libflandmark_static.a
import "C"

import (
	"runtime"
	"unsafe"
)

// Cascade is an OpenCV Haar cascade.
type Cascade struct {
	pointer unsafe.Pointer
}

// LoadCascade loads a Haar cascade from a file path.
func LoadCascade(path string) (*Cascade, error) {
	ptr := C.flandmark_binding_cascade_load(C.CString(path))
	if ptr == nil {
		return nil, ErrCouldNotLoad
	}
	res := &Cascade{ptr}
	runtime.SetFinalizer(res, res.free)
	return res, nil
}

func (c *Cascade) Detect(img *Image, factor float64, minNeighbors int,
	minSize Size, maxSize Size) []Rect {
	if !c.valid() || !img.valid() {
		return nil
	}
	rects := C.flandmark_binding_cascade_detect_objects(c.pointer, img.pointer,
		C.double(factor), C.int(minNeighbors), C.int(minSize.Width),
		C.int(minSize.Height), C.int(maxSize.Width), C.int(maxSize.Height))
	count := int(C.flandmark_binding_rects_count(rects))
	list := make([]Rect, count)
	for i := 0; i < count; i++ {
		var resList [4]C.int
		C.flandmark_binding_rects_get(rects, C.int(i), &resList[0])
		list[i] = Rect{Point{int(resList[0]), int(resList[1])},
			Size{int(resList[2]), int(resList[3])}}
	}
	C.flandmark_binding_rects_free(rects)
	return list
}

func (c *Cascade) free() {
	C.flandmark_binding_cascade_free(c.pointer)
}

func (c *Cascade) valid() bool {
	return c != nil && c.pointer != nil
}

// Image is an OpenCV-compatible image.
type Image struct {
	pointer unsafe.Pointer
}

// NewRGBAImage creates an image from raw RGBA data.
// The image will automatically be freed.
func NewRGBAImage(data []byte, width int, height int) (*Image, error) {
	if len(data) != 4 * width * height {
		return nil, ErrDataSize
	}
	buffer := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	v := C.flandmark_binding_image_rgba(buffer, C.int(width), C.int(height))
	res := &Image{v}
	runtime.SetFinalizer(res, res.free)
	return res, nil
}

// NewGrayImage creates an image from raw grayscale data.
// The image will automatically be freed.
func NewGrayImage(data []byte, width int, height int) (*Image, error) {
	if len(data) != width * height {
		return nil, ErrDataSize
	}
	buffer := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	v := C.flandmark_binding_image_gray(buffer, C.int(width), C.int(height))
	res := &Image{v}
	runtime.SetFinalizer(res, res.free)
	return res, nil
}

func (i *Image) free() {
	C.flandmark_binding_image_free(i.pointer)
}

func (i *Image) valid() bool {
	return i != nil && i.pointer != nil
}

type Point struct {
	X int
	Y int
}

type Rect struct {
	Point Point
	Size  Size
}

type Size struct {
	Width  int
	Height int
}
