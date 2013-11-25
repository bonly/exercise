package main
import "fmt"

type A struct{
	Name string;
};

type B struct{  //struct中的小于(私有数据)在package中可见,在package外部不公开
	Name string;
};

func main(){
	a:=A{};
	a.Print();
	
	var b TZ=3; //基础类型不能用TZ{3} ?
	b.Print();
	
	(*TZ).Print(&b); //通过类来调对象的方法
	
	var c TZ;
	c.Increate(100);
	fmt.Println(c);
}

func (a A)Print(){  //接收者也是值传递,可用*改为指针传递
	fmt.Println("A");
}

type TZ int;
func (a *TZ) Print(){
	fmt.Println(*a);
}

func (a *TZ) Increate(i int){
	*a += TZ(i);
}
