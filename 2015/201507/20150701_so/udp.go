package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("udp..begin")
	defer fmt.Println("udp..end")

	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("add: ", err)
		return
	}

	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println("add: ", err)
		return
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("list: ", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		rn, remAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("recv: ", err)
			return
		} else {
			fmt.Printf("recv %d bytes from %v: %s\n", rn, remAddr, string(buf[:rn]))
			if remAddr.String() == raddr.String() {
				fmt.Printf("it is i want pack\n")
			}
		}

		lnd, err := conn.WriteToUDP([]byte("hello"), raddr)
		if err != nil {
			fmt.Println("write: ", err)
			return
		} else {
			fmt.Println("send: ", lnd)
		}
	}
}
