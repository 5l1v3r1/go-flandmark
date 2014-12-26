package flandmark

// #include "flandmark_binding.h"
// #cgo linux  pkg-config: opencv
// #cgo darwin pkg-config: opencv
// #cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
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
	runtime.SetFinalizer(res, freeCascade)
	return res, nil
}

// Detect uses the Haar cascade to detect objects and returns a list of
// rectangles for those objects.
func (c *Cascade) Detect(img *Image, factor float64, minNeighbors int,
	minSize Size, maxSize Size) ([]Rect, error) {
	if !c.valid() || !img.valid() {
		return nil, ErrBadArgument
	}
	rects := C.flandmark_binding_cascade_detect_objects(c.pointer, img.pointer,
		C.double(factor), C.int(minNeighbors), C.int(minSize.Width),
		C.int(minSize.Height), C.int(maxSize.Width), C.int(maxSize.Height))
	count := int(C.flandmark_binding_rects_count(rects))
	list := make([]Rect, count)
	for i := 0; i < count; i++ {
		var resList [4]C.int
		C.flandmark_binding_rects_get(rects, C.int(i), &resList[0])
		list[i] = NewRectFromC(resList)
	}
	C.flandmark_binding_rects_free(rects)
	return list, nil
}

func (c *Cascade) valid() bool {
	return c != nil && c.pointer != nil
}

func freeCascade(c *Cascade) {
	C.flandmark_binding_cascade_free(c.pointer)
}

// Image is an OpenCV-compatible image.
type Image struct {
	pointer unsafe.Pointer
}

// NewRGBAImage creates an image from raw RGBA data.
// The image will automatically be freed.
func NewRGBAImage(data []byte, width int, height int) (*Image, error) {
	if len(data) != 4*width*height {
		return nil, ErrDataSize
	}
	buffer := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	v := C.flandmark_binding_image_rgba(buffer, C.int(width), C.int(height))
	if v == nil {
		return nil, ErrUnknown
	}
	res := &Image{v}
	runtime.SetFinalizer(res, freeImage)
	return res, nil
}

// NewGrayImage creates an image from raw grayscale data.
// The image will automatically be freed.
func NewGrayImage(data []byte, width int, height int) (*Image, error) {
	if len(data) != width*height {
		return nil, ErrDataSize
	}
	buffer := (*C.uint8_t)(unsafe.Pointer(&data[0]))
	v := C.flandmark_binding_image_gray(buffer, C.int(width), C.int(height))
	if v == nil {
		return nil, ErrUnknown
	}
	res := &Image{v}
	runtime.SetFinalizer(res, freeImage)
	return res, nil
}

func (i *Image) valid() bool {
	return i != nil && i.pointer != nil
}

func freeImage(i *Image) {
	C.flandmark_binding_image_free(i.pointer)
}

// Model represents the FLANDMARK_Model type.
type Model struct {
	pointer unsafe.Pointer
}

// LoadModel loads a model from a path and returns it.
func LoadModel(path string) (*Model, error) {
	ptr := C.flandmark_binding_model_init(C.CString(path))
	if ptr == nil {
		return nil, ErrCouldNotLoad
	}
	res := &Model{ptr}
	runtime.SetFinalizer(res, freeModel)
	return res, nil
}

// Detect detects landmarks within a given box inside an image.
func (m *Model) Detect(img *Image, box Rect) ([]Point, error) {
	if !m.valid() || !img.valid() {
		return nil, ErrBadArgument
	}
	pointCount := int(C.flandmark_binding_model_M(m.pointer))
	data := make([]C.double, pointCount*2)
	boxList := box.CList()
	ret := C.flandmark_binding_model_detect(m.pointer, img.pointer, &data[0],
		&boxList[0])
	if ret == 1 {
		return nil, ErrNormalize
	} else if ret == 2 {
		return nil, ErrDetect
	}

	// Convert the data points
	res := make([]Point, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		res[i] = Point{int(data[i]), int(data[i+1])}
	}
	return res, nil
}

func (m *Model) valid() bool {
	return m != nil && m.pointer != nil
}

func freeModel(m *Model) {
	C.flandmark_binding_model_free(m.pointer)
}

// Point is a two-dimensional integral Euclidean coordinate.
type Point struct {
	X int
	Y int
}

// Rect is a two-dimensional integral rectangle.
type Rect struct {
	Point Point
	Size  Size
}

// NewRectFromC creates a rectangle from a C array [x, y, width, height].
func NewRectFromC(l [4]C.int) Rect {
	return Rect{Point{int(l[0]), int(l[1])}, Size{int(l[2]), int(l[3])}}
}

// CList creates a C array of the form [x, y, width, height].
func (r Rect) CList() [4]C.int {
	return [4]C.int{C.int(r.Point.X), C.int(r.Point.Y), C.int(r.Size.Width),
		C.int(r.Size.Height)}
}

// Size is a two-dimensional integral rectangular size.
type Size struct {
	Width  int
	Height int
}
