package main

import (
	"fmt"
	"os"
	//"path"
	//"runtime"

	"github.com/lazywei/go-opencv/opencv"
	//"../opencv" // can be used in forks, comment in real application
    "runtime"
    "path"
)
var counter int;

//cached Points are the centerpoints of the markers
var cachedPointsCenterX []int
var cachedPointsCenterY []int

//classifier loaded for object detection
var haarClassifier *opencv.HaarCascade


type BfsHandleItem func(opencv.Point) bool

func main() {
	var edge_threshold int

	win := opencv.NewWindow("Go-OpenCV Webcam")
	defer win.Destroy()

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("can not open camera")
	}
	defer cap.Release()

	win.CreateTrackbar("Thresh", 1, 100, func(pos int, param ...interface{}) {
		edge_threshold = pos
	})

	fmt.Println("Press ESC to quit")
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
                ProcessbetterImage(img, win, edge_threshold)
			} else {
				fmt.Println("Image ins nil")
			}
		}
		key := opencv.WaitKey(10)

		if key == 27 {
			os.Exit(0)
		}
	}
}

func ProcessbetterImage(img *opencv.IplImage, win *opencv.Window, pos int) error {
    _, currentfile, _, _ := runtime.Caller(0)
    //image := opencv.LoadImage(path.Join(path.Dir(currentfile), "../images/eye.jpg"))

    if (counter == 1) {
        //find eye, cache them and then mark them
        counter =0;
        cachedPointsCenterY = []int{}
        cachedPointsCenterX = []int{}

        if (haarClassifier == nil) {
            //load and cache haarClassifier
            haarClassifier = opencv.LoadHaarClassifierCascade(path.Join(path.Dir(currentfile), 
                "20131030_camera_eye.xml"))
        }

        rects := haarClassifier.DetectObjects(img)

        for _, value := range rects {

            var pointRect1 = opencv.Point{value.X() + value.Width(), value.Y()}
            var pointRect2 = opencv.Point{value.X(), value.Y() + value.Height()}
            var centerPointX = int((pointRect1.X + pointRect2.X) / 2)
            var centerPointY = int((pointRect1.Y + pointRect2.Y) / 2)

            var center = opencv.Point{centerPointX, centerPointY}
            cachedPointsCenterX = append(cachedPointsCenterX,center.X)
            cachedPointsCenterY = append(cachedPointsCenterY, center.Y)


            //draw eye marker
            drawEyeMarker(img, center)
        }
    }else {
        //if we have a cached value
        for i := 0; i < len(cachedPointsCenterY); i++ {
            drawEyeMarker(img, opencv.Point{cachedPointsCenterX[i], cachedPointsCenterY[i]})
        }

    }
    win.ShowImage(img)
    //counter for cache
    counter = counter+ 1;
	return nil
}

func drawEyeMarker(img *opencv.IplImage, p opencv.Point) {
    opencv.Circle(img, p, 5, opencv.ScalarAll(175.0), 1, 1, 0)
}

//func calcWhitestColor(field *opencv.IplImage) (*opencv.IplImage) {
//
//    gray := opencv.CreateImage(field.Width(), field.Height(), opencv.IPL_DEPTH_8U, 1)
//    edge := opencv.CreateImage(field.Width(), field.Height(), opencv.IPL_DEPTH_8U, 1)
//    cedge := opencv.CreateImage(field.Width(), field.Width(), opencv.IPL_DEPTH_8U, 3)
//    defer cedge.Release()
//
//    opencv.CvtColor(field, gray, opencv.CV_BGR2GRAY)
//
//    opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
//    opencv.Not(gray, edge)
//
//    // Run the edge detector on grayscale
//    opencv.Canny(gray, edge, float64(50), float64(150*3), 3)
//
//    opencv.Zero(cedge)
//    // copy edge points
//    opencv.Copy(field, cedge, edge)
//    return cedge
//}
//
//
//func getPartOfImage(img *opencv.IplImage, rect *opencv.Rect) *opencv.IplImage {
//    return opencv.Crop(img, rect.X(), rect.Y(), rect.Width(), rect.Height())
//}
