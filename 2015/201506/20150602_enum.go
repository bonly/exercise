package main

import (
	"fmt"
	"reflect"
)

type Enum int;

const (
	A Enum = iota
	B
)

func main(){
	f1();
	// f2();
}

func f1(){
	start := interface{}(A);
	typ := reflect.ValueOf(start).Type();
	
	obj := reflect.New(typ).Elem();
	obj.Set(reflect.ValueOf(B));

	fmt.Printf("%s\n", reflect.TypeOf(obj.Interface()));	
	fmt.Printf("%s\n", reflect.ValueOf(obj).Interface());	


	// fmt.Printf("%v\n", *obj);	
}

func f2(){
	start := interface{}(A);
	typ := reflect.TypeOf(start);

	fmt.Printf("%s\n", typ);

}