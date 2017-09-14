package main

// #cgo pkg-config: opencv
// #cgo darwin pkg-config: opencv
// #cgo freebsd pkg-config: opencv
// #cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
// #include <opencv/cv.h>
// #include <opencv/highgui.h>
import "C"
import "unsafe"

func main() {
	f := C.CString("result.jpg")
	defer C.free(unsafe.Pointer(f))

	gray_img := C.cvLoadImage(f, 0)
	defer C.cvReleaseImage(&gray_img)

	w := gray_img.width
	h := gray_img.height
	size := C.cvSize(C.int(w), C.int(h))
	bin_img := C.cvCreateImage(size, C.IPL_DEPTH_8U, 0)
	defer C.cvReleaseImage(&bin_img)

	C.cvThreshold(unsafe.Pointer(gray_img), unsafe.Pointer(bin_img), 0, 255, C.CV_THRESH_BINARY|C.CV_THRESH_OTSU)

	wn := C.CString("Binary")
	defer C.free(unsafe.Pointer(wn))
	C.cvNamedWindow(wn, 1)
	C.cvShowImage(wn, unsafe.Pointer(bin_img))

	C.cvWaitKey(0)
}