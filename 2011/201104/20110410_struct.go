package main
import "fmt"

type test struct{
};

type person struct{
	Name string;
	Age int;
	Contact struct{  //内部匿名结构,不写type
		Phone, City string;
	};
};

func main(){
	a:=test{};
	fmt.Println(a);
	
	b:=person{};
	fmt.Println(b);
	
	c:=person{};
	c.Name="joe";
	c.Age=19;
	fmt.Println(c);
	
	d:=person{Name:"bonly", Age:21};  //赋值都必须全写变量名?
	d.Contact.Phone="15360534225"; //匿名成员结构值不能写在初始化里!暂还不知方法
	fmt.Println(d);
	
	p := teacher{Name: "joe", Age:19, human: human{Sex:0}};
	s := student{Name: "joe", Age:20, human: human{Sex:1}};
	p.Sex = 100;
	fmt.Println(p);
	fmt.Println(s);
	
  k := B{Name: "A", A:A{Name:"B"}};
  fmt.Println(k.Name, k.A.Name);
}

type human struct{
	Sex int;
};

type teacher struct{
	human;
	Name string;
	Age int;
};

type student struct{
	human;
	Name string;
	Age int;
};

type A struct{
	Name string;
};

type B struct{
	A;
	Name string;
};


