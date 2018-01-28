package main

import "fmt"

type Configuration struct {
	Val   string
	Proxy []struct {
		Address string
		Port    string
	}
}

func main() {
	cfg := Configuration{
		Val: "foo",
		Proxy: []struct {
			Address string
			Port    string
		}{
			{Address: "a", Port: "093"},
		},
	}
	fmt.Println(cfg)
}
