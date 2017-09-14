package main
// +build static

/*
//#cgo pkgconfig: --static sqlite3
//#cgo LDFLAGS: -static
//#cgo pkg-config: sqlite3
#cgo LDFLAGS: -lsqlite3
*/
import "C"



//go build --tags static