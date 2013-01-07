package main;
import (
"flag"
"regexp"
"fmt"
);

func main(){
   flag.Parse();
   if flag.NArg() == 1 {
      re, err := regexp.Compile("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})");
      if err != nil {
         fmt.Printf("regex Compile failed\n");
         return;
      }
 
      if re.MatchString(flag.Arg(0)) == true{
         fmt.Printf("%s is a email address.\n", flag.Arg(0));
      }else{
         fmt.Printf("it is not a email address.\n");
      }
  }
}

