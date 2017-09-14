package main

import (
	"fmt"
	"log"
	"net"
	//
	"github.com/gmallard/stompngo"
)

var exampid = "gosend: "

var nmsgs = 5

// Connect to a STOMP 1.1 broker, send some messages and disconnect.
func main() {
	fmt.Println(exampid + "starts ...")

	// Open a net connection
	// n, e := net.Dial("tcp", "127.0.0.1:61613")
	n, e := net.Dial("tcp", "192.168.1.23:61613")
	if e != nil {
		log.Fatalln(e) // Handle this ......
	}
	fmt.Println(exampid + "dial complete ...")

	// Connect to broker
	eh := stompngo.Headers{"login", "guest", "passcode", "guest"}
	conn, e := stompngo.Connect(n, eh)
	if e != nil {
		log.Fatalln("connect: ", e) // Handle this ......
	}
	fmt.Println(exampid + "stomp connect complete ...")

	// Suppress content length here, so JMS will treat this as a 'text' message.
	s := stompngo.Headers{"destination", "/queue/bonly",
		"suppress-content-length", "true"} // send headers, suppress content-length
	t := fmt.Sprint("{\"configItemCode\":\"bonly\",\"logType\":\"Error\"}") ;
	e = conn.Send(s,  t);
	if e != nil{
		log.Println(e);
	}
	// m := exampid + " message: "
	// for i := 1; i <= nmsgs; i++ {
	// 	t := m + fmt.Sprintf("%d", i)
	// 	e := conn.Send(s, t)
	// 	if e != nil {
	// 		log.Fatalln(e) // Handle this ...
	// 	}
	// 	fmt.Println(exampid, "send complete:", t)
	// }

	// Disconnect from the Stomp server
	eh = stompngo.Headers{}
	e = conn.Disconnect(eh)
	if e != nil {
		log.Fatalln(e) // Handle this ......
	}
	fmt.Println(exampid + "stomp disconnect complete ...")
	// Close the network connection
	e = n.Close()
	if e != nil {
		log.Fatalln(e) // Handle this ......
	}
	fmt.Println(exampid + "network close complete ...")

	fmt.Println(exampid + "ends ...")
}