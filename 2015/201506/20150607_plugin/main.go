package main

import (
	"fmt"
	"plugin"
)

func main(){
	pg, _ := plugin.Open("myplugin.so");
	add, _ := pg.Lookup("Add");
	sum := add.(func(int,int)int)(3,4);
	fmt.Printf("sum=%d\n", sum);
}

