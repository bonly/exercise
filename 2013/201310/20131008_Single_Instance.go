package main

import (
    "fmt"
    "log"
    "os"
    "syscall"
    "unsafe"
)

var (
    modkernel32        = syscall.NewLazyDLL("kernel32.dll")
    procCreateMailslot = modkernel32.NewProc("CreateMailslotW")
)

func singleInstance(name string) error {
    ret, _, _ := procCreateMailslot.Call(
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(`\\.\mailslot\`+name))),
        0,
        0,
        0,
    )
    // If the function fails, the return value is INVALID_HANDLE_VALUE.
    if int64(ret) == -1 {
        return fmt.Errorf("instance already exists")
    }
    return nil
}

func main() {
    log.SetFlags(0)
    // pick a unique id here
    err := singleInstance("ea49ee13-7118-4257-a3c2-0c22fc72310d")
    if err != nil {
        log.Fatalln(err)
        return
    }
    log.Println("all good")
    os.Stdin.Read([]byte{0})
}