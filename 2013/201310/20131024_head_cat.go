package main

import(
	// "fmt"
	// "os"

	"github.com/lazywei/go-opencv/opencv"
	"path"
	"runtime"	
)

func main(){
	win := opencv.NewWindow("camer");
	defer win.Destroy();

	cam := opencv.NewCameraCapture(0);
	if cam == nil{
		panic("can not open camera");
	}
	defer cam.Release();

	_, currentfile, _, _ := runtime.Caller(0);

	for {
		if cam.GrabFrame(){
			img := cam.RetrieveFrame(1);
			if img != nil {
				// fmt.Println(path.Join(path.Dir(currentfile)));
				cascade := opencv.LoadHaarClassifierCascade(
					path.Join(path.Dir(currentfile), "20131024_head.xml"));

				faces := cascade.DetectObjects(img);

				for _, value := range faces {
					opencv.Rectangle(img,
						opencv.Point{value.X() + value.Width(), value.Y()},
						opencv.Point{value.X(), value.Y() + value.Height()},
						opencv.ScalarAll(255.0), 1, 1, 0)
				}
				// fmt.Println("retrieve a img");
				win.ShowImage(img);
			}

			//等10,否则来不及显示
			if key := opencv.WaitKey(1); key == 32 {
			// 	os.Exit(0);
			}
		}
	}
	// opencv.WaitKey(0);
}