package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano());
	for i:=1; i<60; i++{
		x := rand.Int() % 100;
		y := rand.Int() % 100;
		if x < y {
			fmt.Printf("%d + %d = %d\t\t", x, y, x+y);
		}else{
			fmt.Printf("%d - %d = %d\t\t", x, y, x-y);
		}
		if i%3 == 0{
			fmt.Println();
		}
	}
	fmt.Println();
}
/*
go build -ldflags "-s" -gcflags "-N -l -g" 20130417.go
*/
