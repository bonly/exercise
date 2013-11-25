package main

import "os"
import "encoding/binary"
import "fmt"

type array_save_t struct{
  U1 uint64;
  U2 uint64;
  U3 uint64;
};


func main(){
  // r实现了io.ReadWriteCloser
  r, err := os.Open(os.Args[1]);
  fmt.Println(os.Args[1]);
  if err != nil {
    panic(err);
  }
  defer r.Close();
  
  var save array_save_t;
  
  // binary.Read takes an io.Reader, which r 实现了的
  if err := binary.Read(r, binary.LittleEndian, &save); err != nil{
    panic(err);
  }
  
  fmt.Println(save);
}
