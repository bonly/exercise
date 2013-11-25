package main
import "fmt"


// 闭包函数,与函数对象有点类似,应该与lambda表达式等价
func main(){
  a:=A;
  a();
  
  f:=closure(10);
  fmt.Println(f(1));
  fmt.Println(f(2));
  
  B();
  C();
  D();
}

func A(){
  fmt.Println("Func A");
  
  fmt.Println("a");
  defer fmt.Println("b");
  defer fmt.Println("c");
  
  for i:=0; i<3; i++{
  	defer func(){  //作参数传进去的是复制值
  		fmt.Println(i); //非传时去的类似于闭包,是引用
  	}();
  }
}

func closure(x int) func(int) int{  //闭包函数
	fmt.Printf("%p\n", &x);
	return func(y int) int{  //匿名函数, 不能作为顶层函数,必须被包在其它函数中
		fmt.Printf("%p\n", &x);
		return x+y;
	}
}

func B(){
	fmt.Println("func B");
}

func C(){
	defer func(){  //如果不recover处理,程序直接退出
		if err:=recover(); err!=nil{
			fmt.Println("recover in c");
		}
	}();
	panic("panic in c"); //defer注册在panic前
}

func D(){
	fmt.Println("func D");
}
