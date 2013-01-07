package main;
import "fmt"
import "strconv"
import . "flag"
import "math"

func main(){
   Parse();
   fmt.Printf("args: %d\n", NArg());
   //minap, _:= strconv.Atoi(Arg(0));
   minap, _:= strconv.ParseFloat(Arg(0),32);
   maxap, _:= strconv.ParseFloat(Arg(1),32);
   defen, _:= strconv.ParseFloat(Arg(2),32);
   hp,    _:= strconv.ParseFloat(Arg(3),32);
   speed, _:= strconv.ParseFloat(Arg(4),32);

   if NArg() == 5 {
      hpower := (minap + maxap)/2.0 * defen * hp / speed;
      power := math.Pow(hpower, 1.0/3.0);
      fmt.Printf("hpower: %f \t power: %f\n", hpower, power);
   } else {
      fmt.Printf("ap_min ap_max defence hp speed\n");
   }
   return;
}
   
