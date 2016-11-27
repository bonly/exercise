package main

import (
        "fmt"
        "path"
        "runtime"
        "time"

        cv "github.com/hybridgroup/go-opencv/opencv"
        "github.com/hybridgroup/gobot"
        "github.com/hybridgroup/gobot/platforms/opencv"
)

func main() {
        _, currentfile, _, _ := runtime.Caller(0);
        // cascade := cv.LoadHaarClassifierCascade(path.Join(path.Dir(currentfile), "haarcascade_frontalface_alt.xml"));
        cascade := path.Join(path.Dir(currentfile), "haarcascade_frontalface_alt.xml");
        fmt.Println(cascade);

        gbot := gobot.NewGobot();

        window := opencv.NewWindowDriver("window");
        camera := opencv.NewCameraDriver("camera", 0);

        work := func() {
                var image *cv.IplImage;

                gobot.On(camera.Event("frame"), func(data interface{}) {
                        image = data.(*cv.IplImage);
                })

                gobot.Every(1000*time.Millisecond, func() {
                        if image != nil {
                                i := image.Clone();
                                faces := opencv.DetectFaces(cascade, i)
                                //opencv.DetectFaces("/usr/share/opencv/haarcascades/haarcascade_eye.xml", i)
                                i = opencv.DrawRectangles(i, faces, 0, 255, 0, 5)
                                window.ShowImage(i)
                        }

                })
        }

        robot := gobot.NewRobot("faceBot",
                []gobot.Connection{},
                []gobot.Device{window, camera},
                work,
        )

        gbot.AddRobot(robot)

        gbot.Start()
}