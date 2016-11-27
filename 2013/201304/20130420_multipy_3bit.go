package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano());
	for i:=0; i<60; i++{
		x := rand.Int() % 99 + 1;
		y := rand.Int() % 9 + 1;
		z := 10 - y;
		fmt.Printf("%d x %d = \t\t\t\t\t%d\n\n",
		   x*10+y, x*10+z, (x*10+y)*(x*10+z));
	}
}