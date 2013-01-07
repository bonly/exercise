package main
import "regexp"
import "bufio"
import "fmt"
import "os"
import "flag"

func grep(re, filename string){
   regex, err := regexp.Compile(re);

   fh, err := os.Open(filename);
   f := bufio.NewReader(fh);
   
   if err != nil {
      return;
   }
   defer fh.Close();

   buf := make([]byte, 1024);
   for {
      buf, _, err = f.ReadLine();
      if err != nil {
         return;
      }
      s := string(buf);
      if regex.MatchString(s){
         fmt.Printf("%s\n", string(buf))
      }
    }
}

func main(){
    flag.Parse();
    if flag.NArg() == 2 {
       grep(flag.Arg(0), flag.Arg(1))
    }else{
       fmt.Printf("Wrong number of arguments.\n")
    }
}

