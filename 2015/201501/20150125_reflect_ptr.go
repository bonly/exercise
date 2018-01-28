package main

import (
"fmt"
"reflect"
)

type A struct{
	Foot string;
};

func (a *A) PrintFoot(){
	fmt.Println("Foot = ", a.Foot);
}

func main(){
	a := &A{Foot: "afoo"};
	val := reflect.Indirect(reflect.ValueOf(a)); //取值对象，再取其地址
	fmt.Println(val.Type().Field(0).Name);
}