package main 

import (
a "a_class"
b "b_class"
c "control"
)

func main(){
	obja := a.A{};
	var objb = b.B{};

	c.FuncC(obja, objb);
}