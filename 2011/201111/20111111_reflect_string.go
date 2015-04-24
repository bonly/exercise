package main

import "fmt"
import "reflect"

type st struct{
};

func (abc *st) Myfn(str string){
	fmt.Println(str);
}

// func Invoke(any interface{}, name string, args... interface{}) {
//   inputs := make([]reflect.Value, len(args));
//   for i, _ := range args {
//       inputs[i] = reflect.ValueOf(args[i]);
//   }
//   reflect.ValueOf(any).MethodByName(name).Call(inputs);
// }

func main(){
	//Invoke(st{}, "Myfn", "abc");

    params := make([]reflect.Value, 1);
    params[0] = reflect.ValueOf("abc");

    reflect.ValueOf(&st{}).MethodByName("Myfn").Call(params);

    //b := &st{};
    //val := reflect.ValueOf(b);
	// val.Elem(). // Go to *B
	// FieldByName("A").     // Find field named A
	// Addr().               // Take its address, since Test has a method receiver of type *A
	// MethodByName("Test"). // Find its method Test
	// Call([]reflect.Value{})
}