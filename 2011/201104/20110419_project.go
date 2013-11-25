package main
import (
 "fmt"
 )

var c chan string;

func main(){
	c = make(chan string);
	go Pingpong();
	for i:=0; i<10; i++{
		c<-fmt.Sprintf("from main: hello, #%d", i);
		fmt.Println(<-c);
	}
}

func Pingpong(){
	i := 0;
  for{
  	fmt.Println(<-c);
  	c<-fmt.Sprintf("from Pingpong: hi, #%d", i);
  	i++;
  }	
}
