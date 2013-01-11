/*
中文判断:regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname"));
英文判断:regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname"));
邮箱:regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email"));
手机号码:regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile"));
15位身份证:regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); 
18位身份证:regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard"));
*/
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

