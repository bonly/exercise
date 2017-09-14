package main

import (
	"log"
	"sync"
	"time"
)

// the function to be run inside a goroutine. It receives a channel on ch, sleeps for t, then sends t on the channel it received
func doStuff(t time.Duration, ch <-chan chan time.Duration) {
	ac := <-ch
	time.Sleep(t)
	ac <- t
}

func main() {
	// create the channel-over-channel type
	sendCh := make(chan chan time.Duration)

	// start up 10 doStuff goroutines
	for i := 0; i < 10; i++ {
		go doStuff(time.Duration(i+1)*time.Second, sendCh)
	}

	// send channels to each doStuff goroutine. doStuff will "ack" by sending its sleep time back
	recvCh := make(chan time.Duration)
	for i := 0; i < 10; i++ {
		sendCh <- recvCh
	}

	// receive on each channel we previously sent. this is where we receive the ack that doStuff sent back above
	var wg sync.WaitGroup // use this to block until all goroutines have received the ack and logged
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dur := <-recvCh
			log.Printf("slept for %s", dur)
		}()
	}
	wg.Wait()
}