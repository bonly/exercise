package main 

/*
#include <stdio.h>

void print(int *param){
	printf("begin\n");
	for (int idx=0; idx < 2; ++idx){
		printf("%d\t", param[idx]);
	}
	printf("\nend\n");
}
*/
import "C"
import (
"unsafe"
// "fmt"
// "reflect"
)

func main(){
	int_arr := []int32{13441, 23};

	C.print((*C.int)(unsafe.Pointer(&int_arr[0])));
}