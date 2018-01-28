package main

import _ "github.com/murlokswarm/windows"
import (
	"fmt"
	"time"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/log"
)

func main() {
	fmt.Println("Windows example")
	app.Run()

	i := 0
	for {
		time.Sleep(time.Second)
		log.Info(i)
		if i == 3 {
			return
		}
		i++
	}
}