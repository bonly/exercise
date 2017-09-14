package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	// "bufio"
	"sync"
)

var Wg sync.WaitGroup;

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	Wg.Add(1);
	go sayhi(conn);
	// go mustCopy(os.Stdout, conn)
	// mustCopy(conn, os.Stdin)
	//var bufs []byte
	//conn.Read(bufs)
	//act(string(bufs))

	Wg.Wait();
}

func sayhi(conn net.Conn){
	defer Wg.Done();

	conn.Write([]byte("exit"));
	// recv, _ := bufio.NewReader(conn).ReadString('\n');
	recv := make([]byte, 1024);
	conn.Read(recv);
	fmt.Println("we got: ", string(recv));
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func act(op string) {
	fmt.Printf("the buf receive is", op)
	switch op {
	case "exit":
		fmt.Printf("you are now leaving")
		os.Exit(1)
	case "ready":
		fmt.Printf("You are ready to play,please waiting for the other players")
	}
}
