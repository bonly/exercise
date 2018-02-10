package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	kcp "github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
)

var key = []byte("testkey")
var salt = []byte("techappen")
var conn kcp.UDPSession

var ch_connect = make(chan int, 1)
var ch_run = make(chan bool)

func init() {
	go func() {
		for {
			select {
			case cnt := <-ch_connect:
				if cnt == 1 {
					fmt.Printf("begin connect\n")
					Connect(&conn, "127.0.0.1:8989")
				} else {
					fmt.Printf("closing connect\n")
				}
				break //此break只能break case
			case run := <-ch_run:
				if run == false {
					fmt.Printf("run == false\n")
				}
				return //断了for
			}
		}
	}()
}

func Connect(pconn *kcp.UDPSession, addr string) int {
	var err error
	pass := pbkdf2.Key(key, salt, 4096, 32, sha1.New)
	block, _ := kcp.NewSalsa20BlockCrypt(pass)

	pconn, err = kcp.DialWithOptions(addr, block, 10, 3)
	if err != nil {
		fmt.Printf("%v\n", err)
		return -1
	}

	fmt.Printf("connect to %s\n", addr)
	return 0
}

func main() {
	var wg sync.WaitGroup

	ch_connect <- 1

	Sig(&wg)

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
				ch_run <- false
				return //需要断for
			default:
				fmt.Printf("get sig: %v\n", sig)
				break
			}
		}
	}()
}
