package main

import "os"
import "encoding/binary"
import "fmt"

type THead struct{
  Flag [8]byte;
  UK1  [14]byte; //可能包括版本号
  Author  [11]byte; 
  Title   [22]byte;
  UK2  [2]byte; //01 00
  FileName [256]byte;
};

type TField struct{
  UD1  [4]byte;
};

type TData struct{
  Head THead;
  Field TField;
};


func main(){
  // r实现了io.ReadWriteCloser
  r, err := os.Open(os.Args[1]);
  fmt.Println(os.Args[1]);
  if err != nil {
    panic(err);
  }
  defer r.Close();
  
  var save TData;
  
  // binary.Read takes an io.Reader, which r 实现了的
  if err := binary.Read(r, binary.LittleEndian, &save); err != nil{
    panic(err);
  }
  
  fmt.Printf("%#v\n",save);
}
