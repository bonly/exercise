package main 

import (
"fmt"
)

func call_me(i_va int, i_vb int)(ret *int){
	ret = new(int);
 	*ret = i_va + i_vb;
 	return ret;
}

func main(){
	sum := call_me(3, 4);
	fmt.Printf("sum: %d\n", *sum);
}
