package main
import "fmt"
import "reflect"

type User struct{
  Id int;
  Name string;
  Age int;
};

func (u User) Hello(){
  fmt.Println("hello world.");
}

func main(){
   u := User{1, "ok", 12};
   Info(u);
}

func Info(o interface{}){
  t:=reflect.TypeOf(o);
  fmt.Println("type:", t.Name());
  
  if k:=t.Kind(); k!=reflect.Struct{ //检查是否结构
  	fmt.Println("not a struct");
  	return;
  }
  	
  v:= reflect.ValueOf(o);
  fmt.Println("Fields:");
  
  for i:=0; i<t.NumField(); i++{ //字段
    f:=t.Field(i);
    val:=v.Field(i).Interface();
    fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val);
  }
  
  for i:=0; i<t.NumMethod(); i++{ //方法
  	m:=t.Method(i);
  	fmt.Printf("%6s: %v\n", m.Name, m.Type);
  }
}
