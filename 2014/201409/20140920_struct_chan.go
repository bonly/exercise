package main

import (
	"log"
	"sync"
	"time"
)

// the struct that we'll pass over a channel to a goroutine running doStuff
type process struct {
	dur time.Duration
	ch  chan time.Duration
}

// the goroutine function. will receive a process struct 'p' on ch, sleep for p.dur, then send p.dur on p.ch
func doStuff(ch <-chan process) {
	proc := <-ch
	time.Sleep(proc.dur)
	proc.ch <- proc.dur
}

func main() {
	// start up the goroutines
	sendCh := make(chan process)
	for i := 0; i < 10; i++ {
		go doStuff(sendCh)
	}

	// store an array of each struct we sent to the goroutines
	processes := make([]process, 10)
	for i := 0; i < 10; i++ {
		dur := time.Duration(i+1) * time.Second
		proc := process{dur: dur, ch: make(chan time.Duration)}
		processes[i] = proc
		sendCh <- proc
	}

	// recieve on each struct's ack channel
	var wg sync.WaitGroup // use this to block until all goroutines have received the ack and logged
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ch <-chan time.Duration) {
			defer wg.Done()
			dur := <-ch
			log.Printf("slept for %s", dur)
		}(processes[i].ch)
	}
	wg.Wait()
}