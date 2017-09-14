package main 

/*
#cgo pkg-config: opencv
#cgo darwin pkg-config: opencv
#cgo freebsd pkg-config: opencv
#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
#cgo CFLAGS: -Wno-error=implicit-function-declaration 
#include <opencv/cv.h>
#include "20140521_scan.c"
*/
import "C"

func main(){
	C.Main();
}