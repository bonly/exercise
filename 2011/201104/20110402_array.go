package main;
import "fmt"

//golang 中array是值类型而不是引用类型,函数参数中是值copy传递
func main(){
  a := [...]int{1,2,3,7:8};
  fmt.Println(a);
  
  var p *[8]int = &a;
  fmt.Println(p);
  
  x,y := 1, 2;
  b:= [...]*int{&x, &y};  //不用={}?!
  fmt.Println(b);
  
  c := [2]int {1,3};
  d := [2]int {1,3};
  e := [2]int {1,2};
  //f := [1]int {1}; //类型不同,不能比效
  
  fmt.Println(c==d);
  fmt.Println(c==e);
}
