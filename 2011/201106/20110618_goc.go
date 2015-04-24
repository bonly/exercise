package main
import "fmt"

func MyFunction(arg1, arg2 int, arg3 string) int64{
    fmt.Print(arg3);
    return int64(arg1+arg2);  
}


func main(){
  fmt.Print(MyFunction(3, 4, "the value is: "));
}

