package main;
import "os/exec"
import "log"

func main(){
  cmd := exec.Command("sleep", "5");
  err := cmd.Start(); // Strat 要Wait来决定结束,是非阻塞的,Run是阻塞的
  if err != nil {
    log.Fatal(err);
  }
  log.Printf("Waiting for command to finish ...");
  err = cmd.Wait();
  log.Printf("Command finished with error: %v", err);  //%v表示value的意思
}
