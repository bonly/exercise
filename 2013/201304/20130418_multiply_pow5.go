package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano());
	for i:=0; i<60; i++{
		x := rand.Int() % 10 * 10 + 5;
		y := x;
		fmt.Printf("%d x %d = \t\t\t\t\t\t\t\t\t%d\n\n", 
		        x, y, x*y);
	}
}