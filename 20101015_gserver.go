package main;
import "os/exec";
import "os";
import "log";

func main(){
  /*
  _, er := exec.LookPath("GServer");
  if er != nil{
     log.Fatal("LookPath: ", er);
  }
  */
  cmd := exec.Command("/home/bonly/workspace/GServer/Debug/GServer", "/home/bonly/workspace/GServer/Debug/bsrv.conf");
  cmd.Stdout = os.Stdout;
  err := cmd.Start(); // Strat 要Wait来决定结束,是非阻塞的,Run是阻塞的
  //err := cmd.Run(); // Strat 要Wait来决定结束,是非阻塞的,Run是阻塞的
  if err != nil {
    log.Fatal(err);
  }
  log.Printf("Waiting for command to finish ...");
  //err = cmd.Wait();
  log.Printf("Command finished with error: %v", err);  //%v表示value的意思
}
