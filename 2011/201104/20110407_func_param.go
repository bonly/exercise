package main;
import "fmt";

func main(){
  a, b:=1,2;
  A(a,b);
  fmt.Println(a,b);
  
  B(a, b);
  fmt.Println(a,b);
  
  s1 := []int{1,2,3,4};
  C(s1);
  fmt.Println(s1);
}

func A(a ... int){
  fmt.Println(a);
}

func B(s ... int){  //不定长变参传的是相当于slice的值复制,值传递是副本
	s[0] = 3;
	s[1] = 4;
	fmt.Println(s);
}

func C(s []int){ //如果真的传slice 相当于传指针(地址值的复制),可修改外部值
	s[0]=5;
	s[1]=6;
	s[2]=7;
	s[3]=8;
	fmt.Println(s);
}
