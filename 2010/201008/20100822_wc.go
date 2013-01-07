package main;
import "fmt";
import "bytes";
import _ "os";
import "strconv";
import "os/exec";

func check(){
  cmd0 := exec.Command("netstat","-a");
  cmd1 := exec.Command("grep", "22");
  cmd2 := exec.Command("wc");
  cmd3 := exec.Command("awk", "{print $1}");

  cmd1.Stdin, _ = cmd0.StdoutPipe();
  cmd2.Stdin, _ = cmd1.StdoutPipe();
  cmd3.Stdin, _ = cmd2.StdoutPipe();
  var output bytes.Buffer;
  //output, _ := cmd3.Output();
  cmd3.Stdout = &output;

  _ = cmd3.Start();
  _ = cmd2.Start();
  _ = cmd1.Start();
  _ = cmd0.Run();
  _ = cmd1.Wait();
  _ = cmd2.Wait();
  _ = cmd3.Wait();

  str := output.String();
  count, err := strconv.ParseInt(str,10, 8);
  if err != nil{
    fmt.Printf("%v\n", err);
  }
  fmt.Printf("count: %d %s\n", count, output.String());
  if count > 0 {
    return;
  }

  //cmd := exec.Command("GServer", "/home/bonly/workspace/GServer/Debug/bsrv.conf");
  //cmd := exec.Command("/bin/ls", "-l");
  //cmd.Stdout = os.Stdout;
  //cmd.Start();
}

func main(){
  for {
    cmd := exec.Command("sleep", "5");
    cmd.Run();
    check();
  }
}

