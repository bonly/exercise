package main

import (
	"fmt"
	"os"

	//"github.com/lazywei/go-opencv/opencv"
	"github.com/lazywei/go-opencv/opencv" // can be used in forks, comment in real application
)

var (
	face_cascade CascadeClassifier
	eyes_cascade CascadeClassifier
	window_name  string = "Capture - Face detection"
)

func main() {
	// PlayVideo("/dos/resource/VID/Tracy_McGrady.avi")
	ShowCamera()
}

func PlayVideo(filename string) {
	cap := opencv.NewFileCapture(filename)
	if cap == nil {
		panic("can not open video")
	}
	defer cap.Release()
	frames := int(cap.GetProperty(opencv.CV_CAP_PROP_FRAME_COUNT))
	win := opencv.NewWindow("GoOpenCV: VideoPlayer")

	win.SetMouseCallback(func(event, x, y, flags int) {

	})

	fmt.Println("Start...")
	for {
		img := cap.QueryFrame()
		if img == nil {
			break
		}

		frame_pos := int(cap.GetProperty(opencv.CV_CAP_PROP_POS_FRAMES))
		if frame_pos >= frames {
			break
		}

		win.ShowImage(img)
		key := opencv.WaitKey(33)
		if key == opencv.CV_EVENT_LBUTTONDOWN {
			os.Exit(0)
		} else if key == 1048608 {
			fmt.Println(key)
			os.Exit(0)
		} else if key == -1 {

		} else {
			fmt.Println("key=", key)
		}
	}
	opencv.WaitKey(0)
}

func ShowCamera() {
	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("can not open camera")
	}
	defer cap.Release()

	win := opencv.NewWindow("GoOpenCV: Camera")

	fmt.Println("Start...")
	for {
		img := cap.QueryFrame()
		if img == nil {
			break
		}

		win.ShowImage(img)
		key := opencv.WaitKey(33)
		if key == opencv.CV_EVENT_LBUTTONDOWN {
			os.Exit(0)
		} else if key == 1048608 {
			fmt.Println(key)
			os.Exit(0)
		} else if key == -1 {

		} else {
			fmt.Println("key=", key)
		}
	}
	opencv.WaitKey(0)
}

func FaceDatection() {
	face_cascade_name := "haarcascade_frontalface_alt.xml"
	eyes_cascade_name := "haarcascade_eye_tree_eyeglasses.xml"

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("can not open camera")
	}
	defer cap.Release()

	win := opencv.NewWindow("GoOpenCV: Camera")

	fmt.Println("Start...")
	for {
		img := cap.QueryFrame()
		if img == nil {
			break
		}

		win.ShowImage(img)
		key := opencv.WaitKey(33)
		if key == opencv.CV_EVENT_LBUTTONDOWN {
			os.Exit(0)
		} else if key == 1048608 {
			fmt.Println(key)
			os.Exit(0)
		} else if key == -1 {

		} else {
			fmt.Println("key=", key)
		}
	}
	opencv.WaitKey(0)
}