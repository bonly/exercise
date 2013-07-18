package main;

import "fmt"
import "os"

var a = 12;
//a int = 12;  //err 全局变量必须用var声名, 不能使用:=
func main(){
   fmt.Fprintf(os.Stderr, " a = %d\n", a);
}
