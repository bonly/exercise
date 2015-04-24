package main

import (
    "syscall"
    "log"
    "unsafe"
)

func main(){   
    h := syscall.MustLoadDLL("kernel32.dll")
    c := h.MustFindProc("GetDiskFreeSpaceExW")
    lpFreeBytesAvailable := int64(0)
    lpTotalNumberOfBytes := int64(0)
    lpTotalNumberOfFreeBytes := int64(0)
    r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("D:"))),
        uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
        uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
        uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
    if r2 != 0 {
        log.Println(r2, err, lpFreeBytesAvailable/1024/1024)
    }
}

/*
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
*/