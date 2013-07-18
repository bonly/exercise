package main
 
import (
    "fmt"
)
 
//#cgo pkg-config: opencv
//#include <cv.h>
//#include <highgui.h>
import "C"
import "unsafe"
 
func main() {
    fmt.Println("Hello World!")
    text := C.CString("Hello World!")
    defer C.free(unsafe.Pointer(text))
    C.cvNamedWindow(text, 1)
    img := unsafe.Pointer(C.cvCreateImage(C.cvSize(640, 480), C.IPL_DEPTH_8U, 1))
    C.cvSet(img, C.cvScalar(0, 0, 0, 0), nil)
    var font C.CvFont
    C.cvInitFont(&font, C.CV_FONT_HERSHEY_SIMPLEX|C.CV_FONT_ITALIC,
        1.0, 1.0, 0, 1, 8)
    C.cvPutText(img, text, C.cvPoint(200, 400), &font,
        C.cvScalar(255, 255, 0, 0))
    C.cvShowImage(text, img)
    C.cvWaitKey(0)
}
