# go-flandmark

This is a Go binding for [flandmark](http://cmp.felk.cvut.cz/~uricamic/flandmark/). The flandmark library detects facial landmarks&mdash;the mouth, eyes, and nose&mdash;in images.

# Installation

Installing **go-flandmark** is not as involved as you might fear. It really only takes a few commands:

    go get github.com/unixpickle/go-flandmark
    cd $GOPATH/github.com/unixpickle/go-flandmark
    make

Assuming you already have OpenCV installed and are connected to the internet, the above should work perfectly.

# Usage

The binding itself couldn't be simpler. You can see the [full documentation](http://godoc.org/github.com/unixpickle/go-flandmark) for more details, or you can look at the wonderful [demonstrations](demo).

In essence, the API has a few types that you will need to work with. These types are as follows:

 * `Cascade` is an OpenCV Haar cascade for face detection
 * `Image` is an OpenCV image object
 * `Model` is a flandmark model for recognizing landmarks
 * `Rect` is a basic type for storing rectangles
 * `Point` is a basic type for storing 2D coordinates

To create a `Cascade`, you will pretty much always use `LoadFaceCascade()`. It will load the default frontalface Haar cascade.

To create a `Model`, you will pretty much always use `LoadDefaultModel()`. It will load a pre-trained model for recognizing the eyes, nose, and mouth.

To create an `Image`, you will pretty much always use `GoGrayImage(image.Image)` to convert Go image types to `Image` objects.

Once you have all of these objects created, you will use the `Cascade` to iterate through faces and the `Model` to find landmarks in those faces.
