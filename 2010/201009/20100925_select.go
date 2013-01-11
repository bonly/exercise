/*
如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。
虽然goroutine是并发执行的，但是它们并不是并行运行的。如果不告诉Go额外的东西，同一时刻只会有一个goroutine执行逻辑代码。利用runtime.GOMAXPROCS(n)可以设置goroutine并行执行的数量。GOMAXPROCS 设置了同时运行逻辑代码的系统线程 的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。以后Go的新版本中调度得到改进后，这将被移除。
*/

package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 1, 1
    for {
        select {
        case c <- x:
            x, y = y, x + y
        case <-quit:
        fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}

/*
在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
select {
case i := <-c:
    // use i
default:
    // 当c阻塞的时候执行这里
}
*/
