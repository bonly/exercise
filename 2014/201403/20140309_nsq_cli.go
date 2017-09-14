package main

import (
  "log"
  "github.com/bitly/go-nsq"
)

func main() {
  config := nsq.NewConfig()
 // w, _ := nsq.NewProducer("127.0.0.1:4150", config)
  w, _ := nsq.NewProducer("192.168.1.10:4150", config)

  err := w.Publish("write_test", []byte("test"))
  if err != nil {
      log.Panic("Could not connect")
  }

  w.Stop()
}