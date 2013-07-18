package main;
import "fmt"

const (
  a = 3;
  b = iota;
  c 
  d
);

func main(){
   fmt.Println(a);
   fmt.Println(b);
   fmt.Println(c);
   fmt.Println(d);
   fmt.Println(^1);// 按位取反
}
