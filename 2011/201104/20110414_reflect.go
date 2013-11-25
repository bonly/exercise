package main
import "fmt"
import "reflect"

type User struct{
  Id int;
  Name string;
  Age int;
};

type Manager struct{
  User;
  title string;
};

func main(){
  m:=Manager{User:User{1,"ok",12}, title:"123"};
  t:=reflect.TypeOf(m);
  
  fmt.Printf("%#v\n", t.Field(0)); //User字段
  fmt.Printf("%#v\n", t.Field(1));//title字段
  
  fmt.Printf("%#v\n", t.FieldByIndex([]int{0,0})); //User->Id
  
  x:=123;
  v:=reflect.ValueOf(&x);  //反射的函数一般需要指针为参数
  v.Elem().SetInt(999);
  fmt.Println(x);
  
  u:=User{1, "ok", 12};
  Set(&u);
  fmt.Println(u);
  
  main1();
}

func Set(o interface{}){
	v:=reflect.ValueOf(o);
	if v.Kind()==reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("can't modi");
		return;
	}else{
		v=v.Elem();
	}
	
	if f:=v.FieldByName("Name"); f.Kind()==reflect.String{ //不确定必找到
		f.SetString("bye");
	}
	
	fn := v.FieldByName("Name1");
	if !fn.IsValid(){
		fmt.Println("no this field");
		return;
	}
	
}

func (u User)Hello(name string){
	fmt.Println("hello", name, ",my name is ", u.Name);
}

func main1(){
	u:=User{1, "ok", 12};
	v:=reflect.ValueOf(u);
	mv := v.MethodByName("Hello");
	
	args:=[]reflect.Value{reflect.ValueOf("joe")};//参数要求是slice
	mv.Call(args); //通过反射动态调用
}