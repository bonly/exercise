package main
import "bufio"
import "os"
import "encoding/binary"

type MyType struct{
  U1 uint64;
  U2 uint64;
  U3 uint64;
};

func main(){
  f, _ := os.Create("test.dat");
  defer f.Close();
  
  w := bufio.NewWriter(f);
  defer w.Flush();
  var tp MyType;
  tp.U1=1;
  tp.U2=2;
  tp.U3=3;
  //for i:=0; i<10; i++{ 
    binary.Write(w, binary.LittleEndian, &tp);
  //}
}
