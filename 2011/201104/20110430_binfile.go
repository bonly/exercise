package main

import (
  "fmt"
  "os"
)

func main(){
  f, err := os.OpenFile("file2.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND,
                         os.ModePerm);
  if err != nil {
    panic(err);
  }
  defer f.Close();
  
  wint, err := f.WriteString("helloworld");
  if err != nil {
    panic(err);
  }
  
  fmt.Printf("%d\n", wint);
  
  _, err = f.Seek(0, 0);
  if err != nil{
    panic(err);
  }
  
  bs := make([]byte, 100);
  rint, err := f.Read(bs);
  if err != nil{
    panic(err);
  }
  
  fmt.Printf("%d, %s\n", rint, bs);
  
}
