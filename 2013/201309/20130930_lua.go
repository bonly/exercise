package main

import (
"github.com/Shopify/go-lua"
"fmt"
)

var global string=`
atext = "abc";
`;

var str string=`
print("hello world");
print(atext);
`;


func main() {
	lu := lua.NewState();
	lua.OpenLibraries(lu);	

	go prt();
	go prt();
	go prt();
	for{
		fmt.Println("in prt main");
		lua.DoString(lu, global);
		if err := lua.DoString(lu, str); err != nil {
			panic(err);
		}			
	}
}


func prt(){
	lu := lua.NewState();
	lua.OpenLibraries(lu);	
	for{
		fmt.Println("in prt");
		lua.DoString(lu, global);
		if err := lua.DoString(lu, str); err != nil {
			panic(err);
		}	
	}
}
/*
证明state非线程安全的，
所有lua都使用同一个,
但是分开跟线程走的，安全的
*/