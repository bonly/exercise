package main

import (
	"fmt"
	"github.com/gmallard/stompngo"
	"log"
	"net"
)

var h, p, dest = "127.0.0.1", "61613", "/topic/haha"

func main() {

	n, e := net.Dial("tcp", net.JoinHostPort(h, p))
	if e != nil {
		log.Fatalln(e)
	}
	ch := stompngo.Headers{"accept-version", "1.2", "host", "localhost", "login", "gooduser", "passcode", "guest"}
	conn, e := stompngo.Connect(n, ch)
	if e != nil {
		log.Fatalln(e)
	}

	u := stompngo.Uuid()
	s := stompngo.Headers{"destination", dest, "id", u}

	r, e := conn.Subscribe(s)
	if e != nil {
		log.Fatalln(e) // Handle this ...
	}

	// Read data from the returned channel
	for i := 1; i <= 10000; i++ {
		m := <-r
		if m.Error != nil {
			log.Fatalln(m.Error)
		}
		//
		fmt.Printf("Frame Type: %s\n", m.Message.Command)
		if m.Message.Command != stompngo.MESSAGE {
			log.Fatalln(m)
		}
		h := m.Message.Headers
		for j := 0; j < len(h)-1; j += 2 {
			fmt.Printf("Header: %s:%s\n", h[j], h[j+1])
		}
		fmt.Printf("Payload: %s\n", string(m.Message.Body))

		if h.Value("subscription") != u {
			fmt.Printf("Error condition, expected [%s], received [%s]\n", u, h.Value("subscription"))
			log.Fatalln("Bad subscription header")
		}
	}

	e = conn.Unsubscribe(s)
	if e != nil {
		log.Fatalln(e)
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