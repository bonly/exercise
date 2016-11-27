package control

import (
"fmt"
)

type InfA interface{
	FuncA();
};

type InfB interface{
	FuncB();
};

func FuncC(a InfA, b InfB){
	a.FuncA();
	b.FuncB();
	fmt.Println("Control");
}