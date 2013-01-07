package main
import "fmt"
import "regexp"

func main(){
  r, err := regexp.Compile("Hello");

  if err != nil {
     fmt.Printf("there is a problem with you regexp.\n");
     return;
  }
  
  if r.MatchString("Hello regular expression.") == true{
     fmt.Printf("Match");
  }else{
     fmt.Printf("no match");
  }
}
