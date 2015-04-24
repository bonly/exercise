package main

import (
    "fmt"
    "flag"
    _ "encoding/json"
    "net"
)
import _ "bytes"

var url *string = flag.String("h", "http://www.lvbody.com:8088", "Host url.");
var times *int = flag.Int("t", 1, "run times.");
var cmd *int = flag.Int("c", 1, "cmd.");

func main() {
  flag.Parse();
  
  for i:=0; i<*times; i++{
    switch(*cmd){
      case 1:
         test_body();
         break;
      default:
         break;
    }
  }
}


type THead struct{
  Flag [8]byte; //00000000
  //Len  [32]byte; //
  Len  int; //
};

func test_body(){
  tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:3001");
  if err!=nil {
    fmt.Println(err.Error);
    panic(err.Error());
  }
  
  conn, err := net.DialTCP("tcp", nil, tcpAddr);
  if err != nil{
    fmt.Println(err.Error);
    panic(err.Error());
  }
  defer conn.Close();
  
  //testHead := THead{Flag:[8]byte("00000000"),Len: 32};
  //copy(testHead.Flag[:], "00000000");
/*
str := "abc"
for k, v := range []byte(str) {
  arr[k] = byte(v)
}
*/

  //var testHead THead;
  a := "this is a test";
  buf := new([]byte);
  _, err = conn.Write(buf);  
}
