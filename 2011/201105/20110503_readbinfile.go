package main

import "os"
import "encoding/binary"

type array_entity_t struct{
  Id uint8;
  Sn [256]byte
};

type array_save_t struct{
  Count uint8;
  Array [64]array_entity_t;
  Checksum uint8;
};

func main(){
  // r实现了io.ReadWriteCloser
  r, err := os.Open(os.Args[0]);
  if err != nil {
    panic(err);
  }
  defer r.Close();
  
  var save array_save_t;
  
  // binary.Read takes an io.Reader, which r 实现了的
  if err := binary.Read(r, binary.LittleEndian, &save); err != nil{
    panic(err);
  }
}
