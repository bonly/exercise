package main

import (
        "time"
        "fmt"

        "github.com/hybridgroup/gobot"
)

func main() {
        gbot := gobot.NewGobot()

        work := func() {
                gobot.Every(1*time.Second, func() {
                        fmt.Println("Hello, human!")
                })
        }

        robot := gobot.NewRobot("drone",
                []gobot.Connection{},
                []gobot.Device{},
                work,
        )
        gbot.AddRobot(robot)

        gbot.Start()
}