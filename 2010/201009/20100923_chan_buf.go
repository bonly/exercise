package main

import "fmt"

func main() {
    c := make(chan int, 2)//修改2为1就报错，修改2为3可以正常运行
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)
}

/*
ch := make(chan type, value)

value == 0 ! 无缓冲（阻塞）
value > 0 ! 缓冲（非阻塞，直到value 个元素）
*/
