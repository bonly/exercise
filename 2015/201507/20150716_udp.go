package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("begin")
	defer fmt.Println("end")

	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9999})
	if err != nil {
		fmt.Println("listening: ", err)
		return
	}
	defer conn.Close()

	var data = make([]byte, 1024)
	for {
		ln, peer, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("recv: ", err.Error())
			return
		}
		if ln <= 0 {
			fmt.Println("recv len 0")
			continue
		}
		fmt.Printf("[%v]: %v\n", peer, string(data[:ln]))

		_, err = conn.WriteToUDP(data, peer)
		if err != nil {
			fmt.Println("write: ", err.Error())
			return
		}
	}
}
