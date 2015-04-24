package main

import "fmt"
import "reflect"
import "encoding/json"

type st struct{
}

var jstr string = `{"func":"Myfu"}`;

type stfn struct{
    Fn string `json:"func"`;
}

func (this *st)Myfu(){
    fmt.Println("echo()");
}

func main() {
    s2 := stfn{};
    json.Unmarshal([]byte(jstr), &s2);

    s := &st{};
    v := reflect.ValueOf(s);
  
    v.MethodByName(s2.Fn).Call(nil);
}
