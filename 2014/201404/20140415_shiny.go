package main

import (
    "image"
    "image/color"
    "log"

    "golang.org/x/exp/shiny/driver"
    "golang.org/x/exp/shiny/screen"
    "golang.org/x/mobile/event/key"
    "golang.org/x/mobile/event/lifecycle"
    "golang.org/x/mobile/event/mouse"
    "golang.org/x/mobile/event/paint"
    "golang.org/x/mobile/event/size"
)

const (
    windowWidth  = 640
    windowHeight = 480
    
    modeNone = 0
    modeDrug = 1
)

var (
    red   = color.RGBA{0x7f, 0x00, 0x00, 0x7f}
    white = color.RGBA{0xff, 0xff, 0xff, 0xff}
    
    r = 20
    x = (windowWidth-r)/2
    y = (windowHeight-r)/2
    
    mode = modeNone
    x1 float32 = 0
    y1 float32 = 0
    x2 float32 = 0
    y2 float32 = 0
)

func main() {
    driver.Main(func(s screen.Screen) {
        w, err := s.NewWindow(
            &screen.NewWindowOptions{
                Width:  windowWidth,
                Height: windowHeight,
            },
        )
        if err != nil {
            log.Fatal(err)
        }
        defer w.Release()

        var sz size.Event
        for {
            e := w.NextEvent()
            switch e := e.(type) {
            case lifecycle.Event:
                if e.To == lifecycle.StageDead {
                    return
                }

            case key.Event:
                if e.Code == key.CodeEscape {
                    return
                }

            case mouse.Event:
                if e.Direction == mouse.DirPress {
                    if x <= (int)(e.X) && (int)(e.X) <= x+r && y <= (int)(e.Y) && (int)(e.Y) <= y+r {
                        mode = modeDrug
                        x1 = e.X
                        y1 = e.Y
                    }
                } else if e.Direction == mouse.DirRelease {
                    mode = modeNone
                } else if mode == modeDrug {
                    x2 = e.X
                    y2 = e.Y
                    x += (int)(x2-x1)
                    y += (int)(y2-y1)
                    x1 = x2
                    y1 = y2
                }
                w.Fill(sz.Bounds(), white, screen.Src)
                w.Fill(image.Rect(x, y, x+r, y+r), red, screen.Over)
                w.Publish()

            case paint.Event:
                w.Fill(sz.Bounds(), white, screen.Src)
                w.Fill(image.Rect(x, y, x+r, y+r), red, screen.Over)
                w.Publish()

            case size.Event:
                sz = e

            case error:
                log.Print(e)
            }
        }
    })
}