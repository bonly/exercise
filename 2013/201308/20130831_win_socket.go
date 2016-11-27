package main

import (
    "bytes"
    "log"
    "net"
    "syscall"
    "unsafe"
)

func main() {
    // WSA開始
    var wsadata syscall.WSAData
    if err := syscall.WSAStartup(0x202, &wsadata); err != nil {
        log.Fatal(err)
    }
    // ソケット取得
    sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_IP)
    if err != nil {
        log.Fatal(err)
    }

    // アドレスの取得
    ip := net.IPv4(0, 0, 0, 0).To4()
    if list, err := net.InterfaceAddrs(); err == nil {
        for _, it := range list {
            log.Print(it)
            ipslice := net.ParseIP(it.String()).To4()
            if !bytes.Equal(ipslice, ip) {
                ip = ipslice
                break
            }
        }
    }

    // bind
    sock_in := syscall.SockaddrInet4{Port: 0}
    copy(sock_in.Addr[:], ip)
    if err := syscall.Bind(sock, &sock_in); err != nil {
        log.Fatal(err)
    }

    // プロミスキャスモードの設定
    ret := uint32(0)
    rcvall := uint32(0x98000001)
    arg := uint32(1)
    if err := syscall.WSAIoctl(sock, rcvall, (*byte)(unsafe.Pointer(&arg)), 4, nil, 0, &ret, nil, 0); err != nil {
        log.Fatal(err)
    }

    // 受信待ちサーバ起動
    resultsrv := startServer()
    // 受け渡し用データ
    an := &anOp{}
    an.init()

    // 非同期設定
    if _, err := syscall.CreateIoCompletionPort(sock, resultsrv.iocp, 0, 0); err != nil {
        log.Print(err)
    }

    var buf syscall.WSABuf
    bufbuf := make([]byte, 0x10000)
    buf.Len = uint32(len(bufbuf))
    buf.Buf = (*byte)(unsafe.Pointer(&bufbuf[0]))

    for {
        var d, f uint32
        // 受信開始
        err := syscall.WSARecv(sock, &buf, 1, &d, &f, &an.o, nil)

        switch err {
        case nil:
        case syscall.ERROR_IO_PENDING:
            err = nil
        default:
            log.Print(err)
        }

        // IOが完了するまでブロック
        r := <-an.resultc

        if bufbuf[9] == 6 && r.err == nil {
            // TCP
            log.Printf("%#v", bufbuf[:r.qty])
        }
    }
}

// ここから下はほとんどgo1.1のコピペ
type ioResult struct {
    qty uint32
    err error
}
type resultSrv struct {
    iocp syscall.Handle
}
type anOp struct {
    o       syscall.Overlapped
    resultc chan ioResult
}

func (o *anOp) init() {
    o.resultc = make(chan ioResult, 1)
}

func startServer() (resultsrv *resultSrv) {
    resultsrv = new(resultSrv)
    var err error
    resultsrv.iocp, err = syscall.CreateIoCompletionPort(syscall.InvalidHandle, 0, 0, 1)
    if err != nil {
        panic("CreateIoCompletionPort: " + err.Error())
    }
    go resultsrv.Run()
    return
}

func (s *resultSrv) Run() {
    var o *syscall.Overlapped
    var key uint32
    var r ioResult
    for {
        r.err = syscall.GetQueuedCompletionStatus(s.iocp, &(r.qty), &key, &o, syscall.INFINITE)
        switch {
        case r.err == nil:
        case r.err == syscall.Errno(syscall.WAIT_TIMEOUT) && o == nil:
            panic("GetQueuedCompletionStatus timed out")
        case o == nil:
            panic("GetQueuedCompletionStatus failed " + r.err.Error())
        default:
        }
        (*anOp)(unsafe.Pointer(o)).resultc <- r
    }
}