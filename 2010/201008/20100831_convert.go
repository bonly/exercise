package main;
import "fmt";
import "strconv";
import "strings";

func main(){
   str := "34\n";
   myint, err := strconv.Atoi(strings.Trim(str,"\n"));
   if err != nil{
       panic(err);
   }
   fmt.Printf("%d\n",myint);
}

