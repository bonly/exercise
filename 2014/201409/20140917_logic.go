package main

import (
"fmt"
)

func main(){
	arr := []byte{0x02, 0x03, 0x80, 0x91, 0x30, 0x03};

	fmt.Printf("%X\n", arr);
	sum := arr[0];
	for i := 1; i < len(arr); i++{
		sum = sum ^ arr[i];
	}
	fmt.Printf("校验值: %X\n", sum);
}