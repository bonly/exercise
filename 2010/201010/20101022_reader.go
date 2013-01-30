package main;
import "os";
import "bufio";
import "fmt";
import "log";

func main(){
   fh, err := os.Open("20101021_seek.go");
   if err != nil {
      log.Fatal("open fail: ", err);
   }
   rd64, err := fh.Seek(14, 0);
   if err != nil {
      log.Println(err);
   }

   log.Println("seek to: ", rd64);

   rder := bufio.NewReader(fh);
   buf, _, err := rder.ReadLine();
   if err != nil {
       log.Println(err);
   }

   fmt.Printf("data :\n %s\n", buf);
}


