package main

import (
	"con"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var net con.Net

func main() {
	var wg sync.WaitGroup

	Sig(&wg)
	net.Init()
	fmt.Printf("%#v\n", net)
	net.Work(&wg) //消息处理

	go func() {
		net.Ch_Op <- 1 //连接信号

		time.Sleep(1 * time.Second)

		// ret := make(chan []byte)
		// net.Read(ret) // 读取信号

		net.Ch_Op <- 2

		net.Ch_Op <- 3

		// select {
		// case result := <-ret:
		// 	fmt.Printf("main recv: %s\n", string(result))
		// }
	}()
	wg.Wait()
}

func Sig(Wg *sync.WaitGroup) {
	Wg.Add(1)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer Wg.Done()
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINT:
				fmt.Printf("Get sigint\n")
				net.Stop()

				fmt.Printf("%#v\n", net)
				return //需要断for
			default:
				fmt.Printf("get sig: %v\n", sig)
				break
			}
		}
	}()
}
