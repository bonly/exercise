package main;
import (
"strings"
"regexp"
"fmt"
);

func rmsp_regex(){
  s := "    Institute of Experimental Computer Science      ";
  r, _ := regexp.Compile("\\s*(.*)\\s*");
  res := r.FindStringSubmatch(s);
  fmt.Printf("<%v>\n", res[1]); //res=res[0]?(true/false)  res[1]是结果
}

func rmsp(){
   s := "    Institute of Experimental Computer Science      ";
   fmt.Printf("<%v>\n", strings.TrimSpace(s));
}

func main(){
   rmsp_regex();
   fmt.Println("========");
   rmsp();
}
