package main

// #cgo pkg-config: opencv
// #include <opencv/cv.h>
// #include <opencv/highgui.h>
import "C"

import (
    "unsafe"	
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	fl := C.CString("out_blur.jpg");
	defer C.free(unsafe.Pointer(fl));

	gray_img := C.cvLoadImage(fl, 0);
	defer C.cvReleaseImage(&gray_img);

	out := gosseract.Must(gosseract.Params{Src:"/tmp/out_blur.jpg", Languages:"chi_sim"});

	// Using client digits
	// client, _ := gosseract.NewClient();
	// client = client.Digest("/tmp/dig");
	// out, _ := client.Src("/tmp/out_blur.jpg").Out();
	fmt.Println(out);
}
