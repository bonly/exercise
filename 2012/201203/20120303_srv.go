package main 

import (
"net"
"log"
"fmt"
"container/list"
);

type Peer struct{
  Caddr string;
  Saddr string;
};

func main() {
	var lst *list.List;
	lst = list.New();

	listener, err := net.Listen("tcp", "0.0.0.0:8989");
	if err != nil{
		log.Println(fmt.Sprintf("listen failed: %s", err.Error()));
	}
	defer listener.Close();

	log.Println("Listen...");

    for {
    	conn, err := listener.Accept();
    	if err != nil{
    		log.Println(fmt.Sprintf("accept failed: %s", err.Error()));
    		return;
    	}
    	log.Println("get a connect");
    	lst.PushBack(conn);

    	go handleRequest(conn,lst);
    }

}

// Handles incoming requests.
func handleRequest(conn net.Conn, lst *list.List) {
  fmt.Println("laddr=",conn.LocalAddr().String());
  fmt.Println("laddr=",conn.RemoteAddr().String());
  for {
	  // 建缓存
	  buf := make([]byte, 1024);
	  // 读取内容
	  reqLen, err := conn.Read(buf);
	  if err != nil {
	    fmt.Println("Error reading:", err.Error());
	  }
	  log.Println("recv len: ", reqLen);
	  // 回写消息
	  //conn.Write([]byte("Message received."))
	  
      for item := lst.Front(); item != nil; item = item.Next(){
      	item.Value.(net.Conn).Write([]byte(buf));
      }
	 
  }
  conn.Close();
}