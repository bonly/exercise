package main

import (
    "fmt"
    "io"
    "os/exec"
)

func main() {
    cmd := exec.Command("wc")
    stdin, _ := cmd.StdinPipe()
    io.WriteString(stdin, "hoge")
    stdin.Close()
    out, _ := cmd.Output()
    fmt.Printf("結果: %s", out)
}