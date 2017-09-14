package main

import (
        "fmt"
        "github.com/AkihiroSuda/go-netfilter-queue"
        "os"
)

func main() {
        var err error

        nfq, err := netfilter.NewNFQueue(0, 1000, netfilter.NF_DEFAULT_PACKET_SIZE)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        defer nfq.Close()
        packets := nfq.GetPackets()

        for true {
                select {
                case p := <-packets:
                        fmt.Println(p.Packet)
                        p.SetVerdict(netfilter.NF_ACCEPT)
                }
        }
}