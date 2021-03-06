package main

import (
    "fmt"
    "log"
    "os"

    "github.com/lazywei/go-opencv/opencv"
)

func main() {
    filename := "/tmp/dow.jpg"
    if len(os.Args) == 2 {
        filename = os.Args[1]
    }

    image := opencv.LoadImage(filename)
    if image == nil {
        log.Fatalf("LoadImage %s failed!", filename)
    }
    defer image.Release()

    win := opencv.NewWindow("Go-OpenCV")
    defer win.Destroy()

    win.SetMouseCallback(func(event, x, y, flags int) {
        fmt.Printf("event = %d, x = %d, y = %d, flags = %d\n",
            event, x, y, flags,
        )
    })
    win.CreateTrackbar("Thresh", 1, 100, func(pos int) {
        fmt.Printf("pos = %d\n", pos)
    })

    win.ShowImage(image)
    opencv.WaitKey(0)
}