package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "/path.sock [wwwroot]")
		return
	}

	fmt.Println("Unix HTTP server")

	root := "."
	if len(os.Args) > 2 {
		root = os.Args[2]
	}

	os.Remove(os.Args[1])

	server := http.Server{
		Handler: http.FileServer(http.Dir(root)),
	}

	unixListener, err := net.Listen("unix", os.Args[1])
	if err != nil {
		panic(err)
	}
	server.Serve(unixListener)
}