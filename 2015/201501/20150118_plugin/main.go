package main 

import (
"plugin"
"fmt"
)

func main(){
	p, err := plugin.Open("./myplugin.so");
	if err != nil{
		fmt.Printf("Open: %v\n", err);
		return;
	}


	add, err := p.Lookup("Add");
	if err != nil{
		fmt.Printf("Lookup: %v\n", err);
		return;
	}

	sum := add.(func(int, int) int) (1, 2);
	fmt.Printf("sum: %d\n", sum);
}