package main 

import (
"log"
"log/syslog"
"net/http"
"flag"
)


var srv string;
var lg  *syslog.Writer;

func init(){
  flag.StringVar(&srv, "s", "0.0.0.0:9997", "srv addr"); 

  var err error;
 
  // lg, err := syslog.NewLogger(syslog.LOG_ERR, 1);
  lg, err := syslog.Dial(""/*network*/,""/*addr*/,syslog.LOG_INFO|syslog.LOG_LOCAL0, "bonly"); // connection to a log daemon
  if err != nil{
  	log.Fatal("syslog: ", err);
  }
  lg.Info("log to system");
}

func main(){
	lg.Info(srv);
	log.Fatal(http.ListenAndServe(srv, nil));
}