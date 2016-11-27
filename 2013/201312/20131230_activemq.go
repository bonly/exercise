package main

import (
	"fmt"
	"github.com/gmallard/stompngo"
	"log"
	"net"
)

var h, p, dest = "127.0.0.1", "61613", "/queue/haha"

func main() {

	n, e := net.Dial("tcp", net.JoinHostPort(h, p))
	if e != nil {
		log.Fatalln(e)
	}

	ch := stompngo.Headers{"accept-version", "1.1", "host", "localhost", "login", "gooduser", "passcode", "guest"}
	conn, e := stompngo.Connect(n, ch)
	if e != nil {
		log.Fatalln(e)
	}

	s := stompngo.Headers{"destination", dest}
	m := " message: "
	for i := 1; i <= 100; i++ {
		t := m + fmt.Sprintf("%d", i)
		e := conn.Send(s, t)
		if e != nil {
			log.Fatalln(e)
		}
		fmt.Println("send complete:", t)
	}

	e = conn.Disconnect(stompngo.Headers{})
	if e != nil {
		log.Fatalln(e)
	}

	e = n.Close()
	if e != nil {
		log.Fatalln(e)
	}
}