package main
import "fmt"
import "time"

func main(){
   t:=time.Now();
   fmt.Println(t.Format(time.ANSIC));
   fmt.Println(t.Format("Mon Jan _2 15:04:06 2006"));
   fmt.Println(t.Format("Mon Jan _2 16:04:06 2006"));
}
