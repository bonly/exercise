package main;
import "os";
import "os/exec";

func main(){
  cmd0 := exec.Command("ps","-ef");
  cmd1 := exec.Command("grep", "GServer bsrv");
  cmd2 := exec.Command("grep", "-v", "grep");
  cmd3 := exec.Command("awk", "{print $2}");

  cmd1.Stdin, _ = cmd0.StdoutPipe();
  cmd2.Stdin, _ = cmd1.StdoutPipe();
  cmd3.Stdin, _ = cmd2.StdoutPipe();
  cmd3.Stdout = os.Stdout;

  _ = cmd3.Start();
  _ = cmd2.Start();
  _ = cmd1.Start();
  _ = cmd0.Run();
  _ = cmd1.Wait();
  _ = cmd2.Wait();
  _ = cmd3.Wait();
}
