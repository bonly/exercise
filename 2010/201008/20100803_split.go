package main;
import (
"regexp"
"fmt"
"strings"
);

func split_reg() {
   s := "abc,def,ghi";
   r, _ := regexp.Compile("[^,]+");
   res := r.FindAllString(s, -1);
   fmt.Printf("%v", res);
}

func split(){
   s := "abc,def, ghi";
   res := strings.Split(s, ",");
   fmt.Printf("%v", res);
}

func main(){
   split_reg();
   fmt.Printf("\n");
   split();
}

