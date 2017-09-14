package main

import (
    . "fmt"

    "github.com/go-vgo/robotgo"
)

func main() {
  bit_map := robotgo.CaptureScreen(10, 20, 30, 40)
  Println("...", bit_map)

  fx, fy := robotgo.FindBitmap(bit_map)
  Println("FindBitmap------", fx, fy)

  robotgo.SaveBitmap(bit_map, "test.png")
} 