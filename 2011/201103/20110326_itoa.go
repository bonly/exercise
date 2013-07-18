package main;
import "fmt"
import "strconv"

func main(){
   a := 65;
   b := string(a);
   c := strconv.Itoa(a);
   fmt.Println("a=", a);
   fmt.Println("b=", b);
   fmt.Println("c=", c);
}
