package main
import (
  "fmt"
)

func main(){
  c1, c2 := make(chan int), make(chan string);
  o := make(chan bool);
  go func(){
  	a,b := false, false;
    for{
      select {
      	case v,ok := <-c1:
      	  if !ok {
      	  	if !a{
	      	  	o <- true;
	      	  	a = true;
      	    }
      	  	break;
      	  }
      	  fmt.Println("c1", v);
      	case v,ok := <-c2:
      	  if !ok {
      	  	if !b{
	      	  	o <- true;
	      	  	b = true;
	      	  }
      	   	break;
      	  }
      	  fmt.Println("c2", v);
      }
    }
  }();
  
  c1 <- 1;
  c2 <- "hi";
  c1 <- 3;
  c2 <- "hello";
  
  close (c1);
  close (c2); 
  
	for i:=0; i<2; i++{
	  <-o;  
	}
}

