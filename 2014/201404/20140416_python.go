//libadd.go
package main

import "C"

//export add
func add(left, right int) int {
	return left + right
}

func main() {
}

/*
go build -buildmode=c-shared -o libadd.so libadd.go
*/