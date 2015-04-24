package main

/*
type Cmd struct {
    Path string
    Args []string
    Env []string
    Dir string
    Stdin io.Reader
    Stdout io.Writer
    Stderr io.Writer
    ExtraFiles []*os.File
    SysProcAttr *syscall.SysProcAttr
    Process *os.Process
    ProcessState *os.ProcessState
}
*/

import (
    "fmt"
    "os/exec"
)

func main() {
    argv := []string{"-a"}
    c := exec.Command("ls", argv...)
    d, _ := c.Output()
    fmt.Println(string(d)) //因为装的git bash所以可以用ls -a
    /*
     *  .
     *  ..
     *  command.go
     *  lookpath.go
     */
}