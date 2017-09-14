package  main 

import (
"fmt"
)

type Af interface{
	Close();
	Open();
};

type Bc struct{
	St string;
};


func (this *Bc) Open(){
	fmt.Println(this.St + "open");
}

//证明只需要部分实现接口即可,但不能被作为传参
func (this *Bc) Close(){
	fmt.Println(this.St + "close");
}

func Pl(ac Af){
    ac.Open();
}

func main(){
	var bc Bc;
	bc.Open();
	Pl(&bc); //如果被函数传用了,有未实现部分也是不可以的
	// bc.Close(); //未实现部分不能调用
}