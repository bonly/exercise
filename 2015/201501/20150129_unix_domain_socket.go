package main

import "net"
import "fmt"

func echoServer(c net.Conn) {
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }

        data := buf[0:nr]
        fmt.Printf("Received: %v", string(data))
        _, err = c.Write(data)
        if err != nil {
            panic("Write: " + err.Error());
        }
    }
}

func main() {
    l, err := net.Listen("unix", "/tmp/echo.sock")
    if err != nil {
        println("listen error", err)
        return
    }

    for {
        fd, err := l.Accept()
        if err != nil {
            println("accept error", err)
            return
        }

        go echoServer(fd)
    }
}

/*package main

import (
    "io"
    "log"
    "net"
    "time"
)

func reader(r io.Reader) {
    buf := make([]byte, 1024)
    for {
        n, err := r.Read(buf[:])
        if err != nil {
            return
        }
        println("Client got:", string(buf[0:n]))
    }
}

func main() {
    c, err := net.Dial("unix", "/tmp/echo.sock")
    if err != nil {
        panic(err)
    }
    defer c.Close()

    go reader(c)
    for {
        _, err := c.Write([]byte("hi"))
        if err != nil {
            log.Fatal("write error:", err)
            break
        }
        time.Sleep(1e9)
    }
}
*/