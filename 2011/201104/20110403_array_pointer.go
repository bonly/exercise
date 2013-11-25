package main;
import "fmt"

func main(){
  a:=[10]int{};
  a[1]=2;
  fmt.Println(a);
  p:=new ([10]int);
  (*p)[1]=2;  //指针与非指针用相同的方式索引操作数组
  p[2]=3;
  fmt.Println(p);
}
