package main

import "fmt"
/*
int main2();
*/
import "C"

//export HelloFromGo
func HelloFromGo() {
   fmt.Printf("Hello from Go!\n")
}

func main() {
   C.main2()
}
