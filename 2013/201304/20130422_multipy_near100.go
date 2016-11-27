package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano());
	abs_near := 10;
	begin := 100 - abs_near;
	
	for i:=0; i<60; i++{		
		x := rand.Int() % (abs_near*2);
		y := rand.Int() % (abs_near*2);
		
		fmt.Printf("%d x %d = \t\t\t\t%d\n\n",
		   (x+begin), (y+begin), (x+begin)*(y+begin));
	}
}