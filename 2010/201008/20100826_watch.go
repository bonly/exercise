package main;
//import _ "os";
import "os/exec";
import "fmt";
import "bytes";
import "net/http";
//import "encoding/binary";
//import "strconv";
//import "strings";

var connect_count string;

func watch(w http.ResponseWriter, r *http.Request){
     get_count();
     fmt.Fprintf(w, "实时在线人数为: %s", connect_count);
}

func get_count(){
  cmd0 := exec.Command("netstat", "-an");
  cmd1 := exec.Command("grep", "8098");
  cmd2 := exec.Command("wc", "-l");

  cmd1.Stdin, _ = cmd0.StdoutPipe();
  cmd2.Stdin, _ = cmd1.StdoutPipe();
  //cmd2.Stdout = os.Stdout;
  var output bytes.Buffer;
  cmd2.Stdout = &output;

  cmd2.Start();
  cmd1.Start();
  cmd0.Run();
  cmd1.Wait();
  cmd2.Wait();

  connect_count = output.String();
  /*connect_count, err := strconv.ParseInt(strcount, 10, 0);
  //connect_count, err := strconv.Atoi(strcount);
  if err != nil{
     fmt.Println(err);
  }
  */
  //fmt.Printf("%s", connect_count);
}

func main(){
   http.HandleFunc("/", watch);
   http.ListenAndServe(":8090", nil);
}

