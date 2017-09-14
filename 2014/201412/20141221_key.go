package main

import (
"fmt"
"github.com/go-vgo/robotgo"
)

func main(){
	key := robotgo.AddEvent("q");
	if key == 0{
		fmt.Println("press q");
	}
}