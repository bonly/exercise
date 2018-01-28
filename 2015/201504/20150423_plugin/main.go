package main 

import (
	"plugin"
	"fmt"
)

func main(){
	pg, err := plugin.Open("./myplugin.so");
	if err != nil{
		fmt.Printf("open plugin failed: %v\n", err);
		return;
	}

	add, _ := pg.Lookup("Add");
	if err != nil{
		fmt.Printf("func not found: %v\n", err);
		return;
	}

	// sum := add.(func(int,int)int)(3, 4);
	addf := add.(func(int,int)int);
	sum := addf(3,4);
	fmt.Printf("sum: %d\n", sum);
}