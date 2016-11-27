package main 

import "C"

//export Foo
func Foo() int32{
	return 43;
}

func main(){}

/*
go build -buildmode=c-shared -linkshared -ldflags "-r=."
go build -buildmode=c-shared -linkshared 
*/

