package main  

import (
"github.com/lazywei/go-opencv/opencv"
)

func main(){
	cap := NewVideoCapture("");

	win :=  opencv.NewWindow("img");
	defer win.Destroy();

	for{
    	img := cap.GetFrame();
		win.ShowImage(img);
		opencv.WaitKey(10);
	}
}

type VideoCapture struct {
    fileName string
    isCameraMode bool
    capture *opencv.Capture
    Fourcc uint32
    Fps float32
    Size opencv.Size
}

func NewVideoCapture(videoFilePath string) VideoCapture {
    var camMode = false
    if (videoFilePath == "") {
        camMode = true
    }
    var cap *opencv.Capture
    if (videoFilePath == "") {
        cap = opencv.NewCameraCapture(0)
    } else {
        cap = opencv.NewFileCapture(videoFilePath)
    }


    var size = opencv.Size{int(cap.GetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH)), int(cap.GetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT))}
    return VideoCapture{videoFilePath, camMode, cap, uint32(cap.GetProperty(opencv.CV_CAP_PROP_FOURCC)), float32(cap.GetProperty(opencv.CV_CAP_PROP_FPS)), size}
}

func (cap *VideoCapture) GetFrame() *opencv.IplImage {
    if (cap.capture.GrabFrame()) {
        return cap.capture.RetrieveFrame(1)
    }
    return nil
}
