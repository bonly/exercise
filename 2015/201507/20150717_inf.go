package main

import (
	"fmt"
	"sync"
)

type Inf interface {
	Write()
}

type wr struct {
}

func (cp wr) Write() {
	fmt.Println("in write")
}

type Net struct {
	socket int
	sync.Mutex
	wr
}

func main() {
	fmt.Println("begin")
	defer fmt.Println("end")

	var net Net
	net.Lock()
	net.Unlock()
	net.Write()
}
