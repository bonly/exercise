package main 
import (
"fmt"
"reflect"
"unsafe"
)

func main(){
	var c int = 1

	intSlice := []int{100, 1, 2, 3, 4}
	newSlice := intSlice[c:]


	fmt.Printf("Points to the Slice %p\n",&intSlice) //0xc20005d020
	fmt.Printf("Points to the first Element of the underlying Array: %d\n",&intSlice[0]) //833223995872


	//Important!!!!!!!!
	fmt.Printf("Points to the newSlice first Element and not to the Array: %d\n",&newSlice[0]) //833223995880
	fmt.Printf("newSlice[0]: %v\n",newSlice[0]) //0


	ref := reflect.ValueOf(newSlice)
	t := reflect.TypeOf(newSlice)

	//Start address of the underlying Array. But that´s critical in my opinion.
	addr := int(ref.Pointer()) - (t.Align() * c) //833223995872

	fmt.Printf("Addr of the underlying data: %d\n",addr) //833223995872

	underArray := (*[5]int)(unsafe.Pointer(uintptr(addr)))
	fmt.Println( *underArray ) //[100 1 2 3 4]
}