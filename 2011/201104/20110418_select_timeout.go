package main
import "fmt"
import "time"

func main(){
	c:=make(chan bool);
	select {
		case v:=<-c:
		  fmt.Println(v);
		case <- time.After(3*time.Second):
		  fmt.Println("timeout");
	}
}
