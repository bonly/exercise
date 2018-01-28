package main

import "fmt"

type M int

const (
	ABC M = iota
	CDE
)

func main() {
	c := CDE
	fmt.Printf("%v \n", c == CDE)
	fmt.Printf("%v \n", c == ABC)
	ts(c)
}

func ts(in interface{}) {
	nw := in.(M)
	fmt.Printf("%v\n", nw == CDE)
	fmt.Printf("%v\n", nw == ABC)
}
