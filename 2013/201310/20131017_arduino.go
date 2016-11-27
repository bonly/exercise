package main

import (
        // "time"

        "github.com/hybridgroup/gobot"
        "github.com/hybridgroup/gobot/platforms/firmata"
        "github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
        gbot := gobot.NewGobot()

        firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")
        led := gpio.NewLedDriver(firmataAdaptor, "led", "13")
        button := gpio.NewButtonDriver(firmataAdaptor, "button", "2")

        work := func() {
                gobot.On(button.Event("push"), func(data interface{}) {
                        led.On()
                })
                gobot.On(button.Event("release"), func(data interface{}) {
                        led.Off()
                })
        }

        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{led},
                work,
        )

        gbot.AddRobot(robot)

        gbot.Start()
}