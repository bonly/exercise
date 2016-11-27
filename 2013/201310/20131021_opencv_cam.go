package main

import(
	"fmt"
	// "os"

	"github.com/lazywei/go-opencv/opencv"
)

func main(){
	win := opencv.NewWindow("camer");
	defer win.Destroy();

	cam := opencv.NewCameraCapture(0);
	if cam == nil{
		panic("can not open camera");
	}
	defer cam.Release();

	for {
		if cam.GrabFrame(){
			fmt.Println("grab a frame");
			img := cam.RetrieveFrame(1);
			if img != nil {
				fmt.Println("retrieve a img");
				win.ShowImage(img);
			}

			//等10,否则来不及显示
			if key := opencv.WaitKey(10); key == 32 {
			// 	os.Exit(0);
			}
		}
	}
	// opencv.WaitKey(0);
}