package main

import "os"
import "encoding/binary"
import "fmt"

type THead struct{
  Flag [8]byte; //'E','O','F','S','O','N','F','\0'   //F是版本号?
  UK1  [8]byte; //可能包括版本号
  UK2  [4]byte; //修正版本号?
};

type TChart struct{
  Revision [4]byte;  //Project revision number
  Time_Format byte;  //Timing format (0=Milliseconds,1=Deltas)
  Time_Division [4]byte; //Time division 当time_formate==1时才有用
};

type TIniStr struct{
  IniLen  [2]byte; // ini setting length
  IniData byte; // ini setting string
};

/*
Num [2]byte; //Number of INI strings
IniType [2]byte; // ini setting type

Num [2]byte; //Number of Ini boolean
type TIniBool struct{
  IniType byte;
};

Num [2]byte; // Number of Ini integer
type TIniInt struct{
  IniType byte;
  IniData [4]byte;   // data of int
};
*/
type TData struct{
  Head THead;
  Title TIniStr;
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
