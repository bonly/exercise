package main 

import (
"fmt"
"strconv"
)

func main(){
	fmt.Println(string(0x4f55));
	se, _ := strconv.ParseUint("4f55", 16, 32);
	fmt.Println(string(se));
}
