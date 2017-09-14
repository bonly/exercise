package main

import (
    . "fmt"

    "github.com/go-vgo/robotgo"
)

func main() {
  keve := robotgo.LEvent("k")
  if keve == 0 {
    Println("you press...", "k")
  }

  mleft := robotgo.LEvent("mleft")
  if mleft == 0 {
    Println("you press...", "mouse left button")
  }
} 