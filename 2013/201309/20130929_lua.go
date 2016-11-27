package main

import (
"github.com/Shopify/go-lua"
"fmt"
)

var str string=`
print("hello world");
`;

var lu *lua.State;

func main() {
	lu = lua.NewState();
	lua.OpenLibraries(lu);

	go prt(lu);
	go prt(lu);
	go prt(lu);
	for{
		fmt.Println("in prt main");
		if err := lua.DoString(lu, str); err != nil {
			panic(err);
		}			
	}
}


func prt(lu *lua.State){
	for{
		fmt.Println("in prt");
		if err := lua.DoString(lu, str); err != nil {
			panic(err);
		}	
	}
}
/*
证明state非线程安全的，
所有lua都使用同一个
*/