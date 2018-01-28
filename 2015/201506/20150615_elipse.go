package main

import (
	"fmt"
)

func AFun(lst ...int) {
	for idx, val := range lst {
		fmt.Printf("%d: %d\n", idx, val)
	}
}

func main() {
	lst := []int{3, 9, 6, 14}
	AFun(lst...)
}
