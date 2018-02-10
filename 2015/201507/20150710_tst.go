package main

import "fmt"

type abc struct {
	a int
}

func main() {
	tha := &abc{11}
	//tha = nil

	a := (map[bool]interface{}{false: tha, true: &abc{}})[tha == nil].(*abc).a
	fmt.Printf("%v\n", a)
}
