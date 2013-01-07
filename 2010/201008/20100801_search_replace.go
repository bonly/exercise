package main;
import (
"flag"
"regexp"
"bufio"
"fmt"
"os"
)

func replace(re, repl, filename string) {
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
     str := string(buf);
     result := regex.ReplaceAllString(str, repl);
     fmt.Print(result + "\n");
  }
}

func main(){
   flag.Parse();
   if flag.NArg() == 3 {
      replace( flag.Arg(0), flag.Arg(1), flag.Arg(2));
   } else {
      fmt.Printf("Wrong number of arguments. \n");
   }
}

