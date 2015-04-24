package main

import "os"
import "encoding/binary"
import "fmt"
import "bytes"

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
  IniData []byte; // ini setting string
};

type IRead interface{
  read();
};

func (str* TIniStr) read(fl *os.File){
   if err := binary.Read(fl, binary.LittleEndian, &str.IniLen); err != nil{
	    panic(err);
   }
   
   bti := str.IniLen[:]; ///转换为切片
   var ilen uint16;
   _ = binary.Read(bytes.NewReader(bti), binary.LittleEndian, &ilen); ///先转成有read功能的bytes
   
   if ilen > 0 {
	   str.IniData = make([]byte, ilen);
	   fl.Read(str.IniData);
   }
}

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
  
  var save THead;
  
  // binary.Read takes an io.Reader, which r 实现了的
  if err := binary.Read(r, binary.LittleEndian, &save); err != nil{
    panic(err);
  }
  
  fmt.Printf("文件头:%#v\n",save);
  
  var 作者 TIniStr;
  作者.read(r);
  fmt.Printf("作者:%#v\n",作者);
  
  var 歌名 TIniStr;
  歌名.read(r);
  fmt.Printf("歌名:%#v\n",歌名);
  
  var 制作者 TIniStr;
  制作者.read(r);
  fmt.Printf("制作者:%#v\n",制作者);
  
  var 未知1 TIniStr;
  未知1.read(r);
  fmt.Printf("未知1:%#v\n",未知1);
  
  var 标注 TIniStr;
  标注.read(r);
  fmt.Printf("标注:%#v\n",标注);
    
  var wlen [4]byte;
  binary.Read(r, binary.LittleEndian, &wlen);
  fmt.Printf("未知2:%#v\n",wlen);
  
  var 文件名 [10]byte;
  binary.Read(r, binary.LittleEndian, &文件名);
  fmt.Printf("文件名:%#v\n",文件名);  
  
  var 未知3 [6]byte;
  binary.Read(r, binary.LittleEndian, &未知3);
  fmt.Printf("未知3:%#v\n",未知3);  
}
