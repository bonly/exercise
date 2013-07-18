package main
import "fmt"

func main(){
LABEL1:
  for {
     for i:=0; i<10; i++{
        if i > 3 {
           break LABEL1;  // break continue goto 都可后跟label 
        }
     }
  }
  fmt.Println("finish");
}
          