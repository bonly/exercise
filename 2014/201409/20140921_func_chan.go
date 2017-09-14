package main

import (
	"log"
	"sync"
	"time"
)

func doStuff(dur time.Duration, ch <-chan func(time.Duration)) {
	ackFn := <-ch
	time.Sleep(dur)
	ackFn(dur)
}

func main() {
	// start up the doStuff goroutines
	sendCh := make(chan func(time.Duration))
	for i := 0; i < 10; i++ {
		dur := time.Duration(i+1) * time.Second
		go doStuff(dur, sendCh)
	}

	// create the channels that will be closed over, create functions that close over each channel, then send them to the doStuff goroutines
	recvChs := make([]chan time.Duration, 10)
	for i := 0; i < 10; i++ {
		recvCh := make(chan time.Duration)
		recvChs[i] = recvCh
		fn := func(dur time.Duration) {
			recvCh <- dur
		}
		sendCh <- fn
	}

	// receive on the closed-over functions
	var wg sync.WaitGroup // use this to block until all goroutines have received the ack and logged
	for _, recvCh := range recvChs {
		wg.Add(1)
		go func(recvCh <-chan time.Duration) {
			defer wg.Done()
			dur := <-recvCh
			log.Printf("slept for %s", dur)
		}(recvCh)
	}
	wg.Wait()
}