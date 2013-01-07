package main
import (
"fmt"
"strings"
"regexp"
);

func sub_reg(){
  s :="OttoFritzHermanWaldoKarlSiegfried";
  r , _ := regexp.Compile("Waldo");
  res := r.MatchString(s);
  fmt.Printf("%v", res);
}

func sub(){
  s :="OttoFritzHermanWaldoKarlSiegfried";
  res := strings.Index(s, "Waldo");
  fmt.Printf("%v", res != -1);
}

func main(){
  sub_reg();
  fmt.Printf("\n");
  sub();
}
