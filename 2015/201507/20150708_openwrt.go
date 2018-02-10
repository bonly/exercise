package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
GOARCH=mipsle GOMIPS=softfloat go build 20150708_openwrt.go
GOARCH=mips GOMIPS=softfloat go build 20150708_openwrt.go
*/
