package main

import (
  "log"
  "sync"

  "github.com/bitly/go-nsq"
)

func main() {

  wg := &sync.WaitGroup{};
  wg.Add(1);

  config := nsq.NewConfig();
  q, _ := nsq.NewConsumer("wechat.text", "bonly", config);
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      log.Printf("Got a message: %v", string(message.Body));
      wg.Done();
      return nil;
  }));
  // err := q.ConnectToNSQD("127.0.0.1:4150");
  err := q.ConnectToNSQD("120.25.106.243:4150");
  //err := q.ConnectToNSQLookupd("devpay.xbed.com.cn:4161");
  if err != nil {;
      log.Panic("Could not connect");
  }
  wg.Wait();

}