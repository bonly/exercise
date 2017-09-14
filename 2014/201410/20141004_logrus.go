package main 

import (
    log "github.com/Sirupsen/logrus"
    // "os"
);

func init() {
    cust_format := new(log.TextFormatter);
    cust_format.TimestampFormat = "2006-01-02 15:04:05.999999";
    cust_format.FullTimestamp = true;
    log.SetFormatter(cust_format);
}

func main(){
  log.Printf("OK\n");
  log.WithFields(log.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")

  log.WithFields(log.Fields{
    "omg":    true,
    "number": 122,
  }).Warn("The group's number increased tremendously!")


  log.WithFields(log.Fields{
    "omg":    true,
    "number": 100,
  }).Fatal("The ice breaks!")

  // A common pattern is to re-use fields between logging statements by re-using
  // the logrus.Entry returned from WithFields()
  contextLogger := log.WithFields(log.Fields{
    "common": "this is a common field",
    "other": "I also should be logged always",
  })

  contextLogger.Info("I'll be logged with common and other field")
  contextLogger.Info("Me too")

}

/*
import (
lg "github.com/Sirupsen/logrus"
"runtime"
"strings"
"fmt"
)

func init(){
    cust_format := new(lg.TextFormatter);
    cust_format.TimestampFormat = "2006-01-02 15:04:05.999999";
    cust_format.FullTimestamp = true;
    lg.SetFormatter(cust_format);
}

func Locate(fields lg.Fields) lg.Fields{
  _, path, line, ok := runtime.Caller(1);
  if ok {
    file := strings.Split(path, "/");
    fields["M"] = fmt.Sprintf("%s:%d", file, line);
  }
  return fields;
}

func WithFields(fields Fields) *lg.Entry{
  return lg.WithFields(Locate(lg.Fields((lg.Fields)(fields))));
}
*/