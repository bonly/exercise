package con

import (
	"crypto/sha1"
	"fmt"
	"sync"
	"time"

	kcp "github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
)

type Net struct {
	key    []byte
	salt   []byte
	conn   *kcp.UDPSession
	Ch_Op  chan int
	ch_run chan bool
}

func (ths *Net) Init() {
	ths.key = []byte("testkey")
	ths.salt = []byte("techappen")
	ths.Ch_Op = make(chan int, 1)
	ths.ch_run = make(chan bool, 1)
}

func init() {
	fmt.Printf("init Net module\n")
}

func (cp Net) Stop() {
	cp.ch_run <- false
}

func (ths *Net) Work(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case cnt := <-ths.Ch_Op:
				switch cnt {
				case 1:
					fmt.Printf("begin connect\n")
					ths.Connect("127.0.0.1:8989")
					break
				case 0:
					fmt.Printf("closing connect\n")
					break
				case 2:
					fmt.Printf("begin read\n")
					ths.Read()
					break
				case 3:
					fmt.Printf("need to write\n")
					ths.Write([]byte("abc"))
					break
				}
				break //此break只能break case
			case run := <-ths.ch_run:
				if run == false {
					fmt.Printf("run == false\n")
				}
				return //断了for
			}
		}
	}()
}

func (ths *Net) Connect(addr string) int {
	var err error
	pass := pbkdf2.Key(ths.key, ths.salt, 4096, 32, sha1.New)
	block, _ := kcp.NewSalsa20BlockCrypt(pass)

	ths.conn, err = kcp.DialWithOptions(addr, block, 10, 3)
	if err != nil {
		fmt.Printf("%v\n", err)
		return -1
	}

	fmt.Printf("connect to %s\n", addr)
	return 0
}

// func (ths *Net) Read(out chan []byte) {
func (ths *Net) Read() {
	buf := make([]byte, 1024)

	ths.conn.SetDeadline(time.Now().Add(1 * time.Second))
	cnt, err := ths.conn.Read(buf)
	if err != nil {
		fmt.Printf("read: %s\n", err.Error())
		return
	}

	fmt.Printf("recv %d: %s\n", cnt, string(buf))
	// out <- buf[:cnt]
}

func (ths *Net) Write(pk []byte) {
	nc, err := ths.conn.Write(pk)
	if err != nil {
		fmt.Printf("send: %s\n", err.Error())
		return
	}

	fmt.Printf("send %d: %s\n", nc, string(pk))
}
