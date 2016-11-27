package main

import (
	"bytes"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/glog"
	"fmt"
)

var (
	icmpSequence      uint16
	icmpSequenceMutex sync.Mutex
)

func getICMPSequence() uint16 {
	icmpSequenceMutex.Lock()
	defer icmpSequenceMutex.Unlock()
	icmpSequence += 1
	return icmpSequence
}

func probeICMP(target string, w http.ResponseWriter, timeout time.Duration) (success bool) {
	deadline := time.Now().Add(timeout)
	socket, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		glog.Info(fmt.Sprintf("Error listening to socket: %s", err));
		return
	}
	defer socket.Close()

	ip, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		glog.Info(fmt.Sprintf("Error resolving address %s: %s", target, err));
		return
	}

	seq := getICMPSequence()
	pid := os.Getpid() & 0xffff

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: pid, Seq: int(seq),
			Data: []byte("Prometheus Blackbox Exporter"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		glog.Info(fmt.Sprintf("Error marshalling packet for %s: %s", target, err));
		return
	}
	if _, err := socket.WriteTo(wb, ip); err != nil {
		glog.Info(fmt.Sprintf("Error writing to socker for %s: %s", target, err));
		return
	}

	// Reply should be the same except for the message type.
	wm.Type = ipv4.ICMPTypeEchoReply
	wb, err = wm.Marshal(nil)
	if err != nil {
		glog.Info(fmt.Sprintf("Error marshalling packet for %s: %s", target, err));
		return
	}

	rb := make([]byte, 1500)
	if err := socket.SetReadDeadline(deadline); err != nil {
		glog.Info(fmt.Sprintf("Error setting socket deadline for %s: %s", target, err));
		return
	}
	for {
		n, peer, err := socket.ReadFrom(rb)
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				glog.Info(fmt.Sprintf("Timeout reading from socket for %s: %s", target, err));
				return
			}
			glog.Info(fmt.Sprintf("Error reading from socket for %s: %s", target, err));
			continue
		}
		if peer.String() != ip.String() {
			continue
		}
		if bytes.Compare(rb[:n], wb) == 0 {
			success = true
			return
		}
	}
	return
}