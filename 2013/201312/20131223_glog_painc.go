package main

import (
"os"
"fmt"
"strings"
)

func main() {
	os.Setenv("GOTRACEBACK", "crash"); 
	for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0], "=", pair[1])
    }	
	panic("kerboom");
}

/*
% env GOTRACEBACK=0 ./crash 
panic: kerboom
% echo $?
2
*/
