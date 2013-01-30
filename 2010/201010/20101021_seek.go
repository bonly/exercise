package main;
import "os";
//import "io";
import "fmt";
import "log";

func main(){
   fh, err := os.Open("20101021_seek.go");
   if err != nil {
      log.Fatal("open fail: ", err);
   }
   rd64, err := fh.Seek(10, 0);
   if err != nil {
      log.Println(err);
   }

   log.Println("seek to: ", rd64);

   buf := make([]byte, 255);
   rd, err := fh.Read(buf);
   if err != nil {
       log.Println(err);
   }

   fmt.Printf("read %d byte: %s", rd, buf);
}


