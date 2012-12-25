package main
import "fmt"
import "os/exec"

func main() {
   out, err := exec.Command("ls", "-l").Output();
   if err != nil {
   	fmt.Print(err);
   	return;
   }
   fmt.Printf("the data is %s\n", out);
}
