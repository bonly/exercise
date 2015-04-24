package main
import "fmt"
import "os"
import "os/signal"
import "syscall"
func main() {
    sigs := make(chan os.Signal, 1); //建接受signal的chan, buffer=1
    done := make(chan bool, 1); //等待结束的chan, buffer=1

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2);  //定义所接受的信号

	go func() {
        sig := <-sigs; //把signal中的消息转给sig变量
        fmt.Println();
        fmt.Println(sig); //打印收到的消息
        if sig == syscall.SIGUSR2{
        	fmt.Println("get a usr2 signal");
        }
        done <- true;  //修改结果chan
    }();

    fmt.Println("awaiting signal");
    <-done; //在此等待结束信号
    fmt.Println("exiting");
}
