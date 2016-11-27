package main 

import (
"github.com/lazywei/go-opencv/opencv"
"fmt"
"os"
)

func main(){
	win := opencv.NewWindow("video", opencv.CV_WINDOW_AUTOSIZE);
	defer win.Destroy();

	cap := opencv.NewFileCapture("/home/bonly/Downloads/soccer.avi");
	if cap == nil{
		panic("can not open video");
	}
	defer cap.Release();

	frames := int(cap.GetProperty(opencv.CV_CAP_PROP_FRAME_COUNT));

	// win.SetMouseCallback(func(event, x, y, flags int){});

	fmt.Println("start...");
	for {
		frame := cap.QueryFrame();
		if frame == nil{
			break;
		}

		frame_pos := int(cap.GetProperty(opencv.CV_CAP_PROP_POS_FRAMES));
		if frame_pos >= frames{
			break;
		}

		win.ShowImage(frame);

		key := opencv.WaitKey(33);
		if key == opencv.CV_EVENT_LBUTTONDOWN{
			os.Exit(0);
		} else if key == 1048608{
			fmt.Println(key);
			os.Exit(0);
		} else if key == -1{
		} else {
			fmt.Println("key=", key);
		}
	}
	opencv.WaitKey(0);
}