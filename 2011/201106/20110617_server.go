package main 

import (
    "flag"
    "fmt"
    "net"
    "os"
    "log"
    "time"
)
const MAX_PLAYERS = 3;
type user struct{
    conn   net.Conn;
};
var allPlayers [MAX_PLAYERS]*user;

const APP_VERSION = "0.1";

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    listen("127.0.0.1:8989");

	for {
		fmt.Print("wait for end...\n");
		time.Sleep(5*1e9);
	}    
}

func listen(addr string) error{
   listener, err := net.Listen("tcp", addr);
   if err != nil{
     return err;
   }
   
   go func(){
     for failures := 0; failures < 3; {
       conn, err := listener.Accept();
       if err != nil{
         log.Print("Failed listenning: ", err, "\n");
         failures++;
       }
       if ok, index := acceptClient(conn); ok{
         go ManageOneClient(conn, index);
         fmt.Print(index);
       }
     }
     log.Println("Too many listener.Accept() errors, giving up");
     os.Exit(1);
   }();
   return nil;
}

func acceptClient(conn net.Conn) (ok bool, index int){
  var i int;
  for i=0; i<MAX_PLAYERS; i++ {
    if allPlayers[i] == nil{
      break;
    }
  }
  if i == MAX_PLAYERS {
    return false, 0;
  }
  up := new(user);
  up.conn = conn;
  allPlayers[i]=up;
  return true, i;
}

func ManageOneClient(conn net.Conn, i int) {
	log.Print("RemoteAddr ", conn.RemoteAddr(), "\n")
	conn.Close();
}

  