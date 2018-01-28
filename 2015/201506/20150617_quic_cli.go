package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"time"

	"github.com/lucas-clemente/quic-go"
)

const addr = "localhost:9999"
const message = "ccc"

func main() {
	session, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	stream, err := session.OpenStreamSync()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Printf("Client: Sending '%s'/n", message)
		_, err = stream.Write([]byte(message))
		if err != nil {
			fmt.Println(err)
			return
		}
		buf := make([]byte, len(message))
		_, err = io.ReadFull(stream, buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Client: Got '%s'/n", buf)
		time.Sleep(2 * time.Second)
	}
}
