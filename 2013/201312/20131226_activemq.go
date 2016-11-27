package main

import (
	"flag"
	"fmt"
	"github.com/guotie/stomp"
	"os"
)

const defaultPort = ":61613";

var serverAddr = flag.String("server", "127.0.0.1:61613", "STOMP server endpoint");
// var serverAddr = flag.String("server", "192.168.1.23:61616", "STOMP server endpoint");
var messageCount = flag.Int("count", 10, "Number of messages to send/receive");
var queueName = flag.String("queue", "/queue/client_test", "Destination queue");
var helpFlag = flag.Bool("help", false, "Print help text");
var stop = make(chan bool);

var options = stomp.Options{
	Login:    "guest",
	Passcode: "guest",
	Host:     "/",
};

func main() {
	flag.Parse();
	if *helpFlag {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0]);
		flag.PrintDefaults();
		os.Exit(1);
	}

	// subscribed := make(chan bool);
	//go recvMessages(subscribed);

	// 等待订阅完成
	//<-subscribed;

	go sendMessages();

	<-stop;
	// <-stop;
}

func sendMessages() {
	defer func() {
		stop <- true;
	}();

	conn, err := stomp.Dial("tcp", *serverAddr, options);
	if err != nil {
		println("cannot connect to server", err.Error());
		return;
	}

	for i := 1; i <= *messageCount; i++ {
		// text := fmt.Sprintf("Message #%d", i);
		text := fmt.Sprintf("aaaafffffffff", i);
		// err = conn.Send(*queueName, "text/plain",
			// []byte(text), nil);
		err = conn.Send(*queueName, "text",
			[]byte(text), nil);
		if err != nil {
			println("failed to send to server", err);
			return;
		}
	}
	println("sender finished");
}

func recvMessages(subscribed chan bool) {
	defer func() {
		stop <- true
	}();

	conn, err := stomp.Dial("tcp", *serverAddr, options);
	if err != nil {
		println("cannot connect to server", err.Error());
		return;
	}

	sub, err := conn.Subscribe(*queueName, stomp.AckAuto);
	if err != nil {
		println("cannot subscribe to", *queueName, err.Error());
		return;
	}
	close(subscribed);

	for i := 1; i <= *messageCount; i++ {
		msg := <-sub.C;
		expectedText := fmt.Sprintf("Message #%d", i);
		actualText := string(msg.Body);
		if expectedText != actualText {
			println("Expected:", expectedText);
			println("Actual:", actualText);
		}
	}
	println("receiver finished");

}