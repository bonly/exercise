package main 

import (
"github.com/lazywei/go-opencv/opencv"
)

func main(){
	var img *opencv.IplImage = opencv.LoadImage("face.jpg",opencv.CV_LOAD_IMAGE_ANYCOLOR);
	win := opencv.NewWindow("image");
	defer win.Destroy();

	for{
		win.ShowImage(img);
		opencv.WaitKey(10);		
	}
}