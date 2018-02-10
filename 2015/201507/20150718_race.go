package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup

	/*
		wg.Add(1)
		go func() {
			defer wg.Done()
			Deposit(200)
			fmt.Println("=", Balance())
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			Deposit(100)
		}()

	*/

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch_dep(200)
		fmt.Println("=", ch_bal())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch_bal()
	}()
	wg.Wait()

}

var ch_deposit = make(chan int)
var ch_balance = make(chan int)

func ch_dep(amount int) {
	ch_deposit <- amount
}

func ch_bal() int {
	return <-ch_balance
}

func op() {
	for {
		select {
		case amount := <-ch_deposit:
			balance += amount
		case ch_balance <- balance:
		}
	}
}

func init() {
	go op()
}
