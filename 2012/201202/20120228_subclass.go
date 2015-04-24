package main 

import (
  "fmt"
)

type A struct{
	AStr string;
};

func (this *A) Print(){
	fmt.Println("A string=", this.AStr);
}

type B struct {
	A;
    AStr string;
}

func (this *B) Print(){
	fmt.Println("B string=", this.AStr);
}

func main(){
	var b B;
    b.A.AStr = "a str";
	b.AStr = "it is a in b";

	b.Print();
	b.A.Print();

	b.AStr = "new str in b";
	b.Print();
	b.A.Print();
}
