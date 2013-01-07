package main;
import "os";
import "os/exec";
import "bytes";
import "log";
import "time";


var connect_count string;

func watch(){
     get_count();
     log.Print(connect_count);
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
   fi, err := os.OpenFile("player_count.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666);
   if err != nil{
      panic(err)
   }
   defer fi.Close();

   log.SetOutput(fi);
   for {
     watch();
     time.Sleep(5 * time.Second);
   }
}

