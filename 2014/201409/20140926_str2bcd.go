package main 

import (
"fmt"
// "strconv"
)

func main(){
	org := "312456";
	var hexd [11]byte;
	for idx:=0; idx<len(org); idx++{
		// chr, err := strconv.Atoi(org[idx]);
		// if err != nil{
		// 	fmt.Printf("conv %s\n", err.Error());
		// }
		hexd[idx] = org[idx] << 4 >> 4;
	}
	fmt.Printf("after: %X\n", hexd);
}
