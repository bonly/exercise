package main;
import "fmt";
import "bytes";
import "os";
import "os/exec";

func check(){
  cmd0 := exec.Command("ps","-ef");
  cmd1 := exec.Command("grep", "GServer bsrv");
  cmd2 := exec.Command("grep", "-v", "grep");
  cmd3 := exec.Command("awk", "{print $2}");

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

  pid_len := len (output.String());
  fmt.Printf("%s %d\n", output.String(), pid_len);
  if pid_len > 0 {
    return;
  }

  cmd := exec.Command("GServer", "/home/bonly/workspace/GServer/Debug/bsrv.conf");
  //cmd := exec.Command("/bin/ls", "-l");
  cmd.Stdout = os.Stdout;
  cmd.Start();
}

func main(){
  for {
    cmd := exec.Command("sleep", "1");
    cmd.Run();
    go check();
  }
}

