package main;
import "fmt"
func main(){
LAB1:
   //k:= 30; //多了这行会报无效的label
   for i:=0; i<10; i++{
      for {
         fmt.Println(i);
         //fmt.Println(k);
         continue LAB1; 
         //goto LAB1;  //无限循化..用goto时,尽量把标签放在goto后面
      }
   }
   fmt.Println("over");
}
      