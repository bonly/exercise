package main
import "fmt"
import "math"
import "strconv"


func main(){
  fmt.Println(strconv.FormatUint(1,2));
  fmt.Println(strconv.FormatUint(math.Float64bits(1.0),2));
  fmt.Println(strconv.FormatUint(1e4, 2));
  fmt.Println(strconv.FormatUint(math.Float64bits(1e4),2));
}

