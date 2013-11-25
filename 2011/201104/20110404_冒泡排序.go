package main;
import "fmt";

func main(){
  a:= [...]int{13,3,45,11,33};
  fmt.Println(a);
  
  num := len(a);
  for i:=0; i<num; i++{
    for j:=i+1; j<num; j++{
      if a[i]<a[j]{
        temp:=a[i];
        a[i] = a[j];
        a[j] = temp;
      }
    }
   }
   fmt.Println(a);
}
