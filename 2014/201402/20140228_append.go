package main

import (
"fmt"
)

func main(){
  ar := []int{1,3,5,7};
  var br []int;
  /* old method
  for _, item := range ar{
    br = append(br, item);
  }
  */
  br = append(br, ar...);  //new method

  for _, item := range br{
    fmt.Println(item);
  }
}
