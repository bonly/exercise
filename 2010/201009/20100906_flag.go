package main;
import "flag";
import "fmt";

var openPort = flag.String("port", ":8080", "http listen port");
var configFile = flag.String("config", "", "config file name");

func main(){
  flag.Parse();
  fmt.Println("Configure filename: ", *configFile);
  fmt.Println("Open Port: ", *openPort);
}
