package main

// /*
//   #include <stdlib.h>
//    typedef void (*Callback)(unsigned int sn, char *buffer);
//    extern Callback fn; //extern的东西需要把实体放到另一个文件中实现
//     extern void SoCallCs(unsigned int sn, char *buffer);
// */

import "C" //要用c-shared必须要加这个

import (
	"fmt"
	"net"
	"runtime/debug"
	"time"
)

func main() {
	// fmt.Println("begin main.") //动态库中，此函数内容被替代了？无效果
	// defer fmt.Println("main end...")

	// Use()
	// Run()

	// time.Sleep(1 * time.Hour)

	// Stop()
}

var run bool = true

// var conn *net.UDPConn
var conn net.Conn
var pc net.PacketConn

// var remote *net.UDPAddr
// var laddr *net.UDPAddr

func init() {
	// debug.SetGCPercent(50)
	// var err error
	// laddr, err = net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	// if err != nil {
	// 	fmt.Println("ladd: ", err)
	// 	return
	// }
	// conn, err = net.ListenUDP("udp", laddr)
	// if err != nil {
	// 	fmt.Println("lis: ", err)
	// 	return
	// }
	// remote, err = net.ResolveUDPAddr("udp", "127.0.0.1:9998")
	// if err != nil {
	// 	fmt.Println("radd: ", err)
	// 	return
	// }

	// conn, err = net.DialUDP("udp", laddr, remote) // 2参(ladd)为nil 是表示随机端口
	// if err != nil {
	// 	fmt.Println("dial: ", err)
	// 	return
	// }
}

func Use() {
	var err error
	pc, err = net.ListenPacket("udp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println(err)
	}
	// defer pc.Close()

	buffer := make([]byte, 1024)
	go func() {
		defer pc.Close()
		for run {
			bnt, addr, err := pc.ReadFrom(buffer)
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Printf("recv %d from %s: %s\n", bnt, addr.String(), string(buffer))
			}
			// pc.WriteTo([]byte("ok"), addr)
			time.Sleep(5 * time.Millisecond)
			debug.FreeOSMemory()
		}
	}()
}

//export Run
func Run() {
	run = true
	var err error
	conn, err = net.Dial("udp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()

	go func() {
		defer conn.Close()
		for run {
			time.Sleep(5 * time.Millisecond)
			debug.FreeOSMemory()
			// fmt.Println("call back to cs")

			// snd, err := conn.WriteToUDP([]byte("hello from so"), remote)
			// if err != nil {
			// 	fmt.Println("write: ", err)
			// } else {
			// 	fmt.Println("send: ", snd)
			// }
			//cbuf := C.CString("abc")
			//defer C.free(unsafe.Pointer(cbuf))
			//C.SoCallCs(C.uint(13), cbuf)
			cnt, err := conn.Write([]byte("hello from so"))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("send ", cnt)
			}
		}
	}()
}

//export Stop
func Stop() int {
	run = false
	return 0
}

// //export SetSoCallCs
// func SetSoCallCs(pt C.Callback) {
// 	C.fn = pt
// }
