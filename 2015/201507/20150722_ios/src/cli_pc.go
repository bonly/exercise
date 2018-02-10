/*
auth: bonly
desc: 跨平台客户端
data: 2018.1.24
*/
package main

import "C"

import (
	"log"
	// "time"
	"net"
	"sync"
)

func init() {}

func main() {
	Srv()
}

//export Srv
func Srv() {
	log.Println("srv begin...")
	defer log.Println("srv end.")

	var Wg sync.WaitGroup

	Local_Srv(&Wg)

	Wg.Wait()
}

func Local_Srv(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if err != nil {
		log.Printf("add: %s\n", err.Error())
		return
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Printf("listen: %s\n", err.Error())
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		ln, peer, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("recv: %s\n", err.Error())
			continue
		}
		log.Printf("qry %d bytes from [%v]: %s\n", ln, peer, string(buf[:ln]))

	}
}
