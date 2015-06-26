package hi

import "fmt"

func Hello(name string) {
	fmt.Println("Hello, %s!\n", name)
}

//go install golang.org/x/mobile/cmd/gobind
//GP gobind -lang=go 20130131_hi > ../go_hi.go
//GP gobind -lang=java 20130131_hi > ../Hi.java
