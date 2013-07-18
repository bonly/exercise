package main;

import "fmt"

func main(){
   switch a:=1; {
     case a>=0:
        fmt.Println("a>=0");
        fallthrough;
     case a>=1:
        fmt.Println("a>=1");
   }
}
