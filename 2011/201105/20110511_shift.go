package main
import (
  "fmt"
)

const (
  ProtVersionMinor = 255;
  ProtVersionMajor = 5;
  CMD_PROT_VERSION = 26;
);

var ClientCurrentMinorVersion = 34;
var ClientCurrentMajorVersion = 5;

func proto(){
  b := []byte{11, 0, CMD_PROT_VERSION,
       ProtVersionMinor & 0xFF,
       ProtVersionMinor >> 8,
       ProtVersionMajor & 0xFF,
       ProtVersionMajor >> 8,
       byte(ClientCurrentMinorVersion & 0xFF),
       byte(ClientCurrentMinorVersion >> 8),
       byte(ClientCurrentMajorVersion & 0xFF),
       byte(ClientCurrentMajorVersion >> 8),
   };
   fmt.Printf("%X\n",b);    
   c := []byte{11, 0, CMD_PROT_VERSION,
	       ProtVersionMinor ,   //在值不大于255时是正确的,大于了就会放不下
	       ProtVersionMinor >> 8,
	       ProtVersionMajor ,
	       ProtVersionMajor >> 8,
	       byte(ClientCurrentMinorVersion),
	       byte(ClientCurrentMinorVersion >> 8),
	       byte(ClientCurrentMajorVersion),
	       byte(ClientCurrentMajorVersion >> 8),
   };   
   fmt.Printf("%X\n",c);    
}

func main(){
  proto();
}
