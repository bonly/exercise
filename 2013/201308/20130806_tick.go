package main

import (
	"fmt"
	"time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR_TO_TICK int = 23
const MINUTE_TO_TICK int = 00
const SECOND_TO_TICK int = 03

func main() {
    ticker := updateTicker()
    for {
	<-ticker.C
	fmt.Println(time.Now(), "- just ticked")
	ticker = updateTicker()
    }
}

func updateTicker() *time.Ticker {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick, "- next tick")
	diff := nextTick.Sub(time.Now())
	return time.NewTicker(diff)
}