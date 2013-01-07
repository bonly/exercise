package main
import "os"
import "bufio"
import "fmt"
import "io"

func main(){
   var file *os.File;
   var err error;
   if file, err = os.Open("foo.txt"); err != nil {
      return
   }
   reader := bufio.NewReader(file);
   for {
       var str string;
       if str, err = reader.ReadString('\n'); err != nil {
           break;
       }
       fmt.Printf("%s\n", str);
    }
    if err == io.EOF{
       err = nil
    }
    return;
}

       
