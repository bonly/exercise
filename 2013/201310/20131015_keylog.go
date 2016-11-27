package main

import (
    "fmt"
    "github.com/MarinX/keylogger"
)

func main() {
    devs, err := keylogger.NewDevices()
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, val := range devs {
        fmt.Println("Id->", val.Id, "Device->", val.Name)
    }

    //keyboard device file, on your system it will be diffrent!
    rd := keylogger.NewKeyLogger(devs[3])

    in, err := rd.Read()
    if err != nil {
        fmt.Println(err)
        return
    }

    for i := range in {

        //we only need keypress
        if i.Type == keylogger.EV_KEY {
            fmt.Println(i.KeyString())
        }
    }
}