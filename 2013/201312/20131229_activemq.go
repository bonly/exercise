package main

import (
  // "os"
  "fmt"
  // "errors"
  "net"
  // "encoding/json"
  "github.com/gmallard/stompngo"
)
/*
type Client struct {
  Host            string
  Port            string
  User            string
  Password        string
  Uuid            string
  NetConnection   net.Conn
  StompConnection *stompngo.Connection
}

func (client *Client) setOpts() {
  client.Host = "localhost"
  client.Port = "61613"
  client.User = ""
  client.Password = ""
  client.Uuid = stompngo.Uuid()
  
  cli_host := os.Getenv("STOMP_HOST")
  if cli_host != "" {
    client.Host = cli_host
  }

  cli_port := os.Getenv("STOMP_PORT")
  if cli_port != "" {
    client.Port = cli_port
  }

  cli_user := os.Getenv("STOMP_USER")
  if cli_user != "" {
    client.User = cli_user
  }

  cli_password := os.Getenv("STOMP_PASSWORD")
  if cli_password != "" {
    client.Password = cli_password
  }

  cli_uuid := os.Getenv("STOMP_UUID")
  if cli_uuid != "" {
    client.Uuid = cli_uuid
  }
}

func (client *Client) netConnection() (conn net.Conn, err error) {
  conn, err = net.Dial("tcp", net.JoinHostPort(client.Host, client.Port))
  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  client.NetConnection = conn
  return
}

func (client *Client) stompConnection() (*stompngo.Connection) {
  headers := stompngo.Headers{
    "accept-version", "1.1", 
    "host", client.Host,
    "login", client.User,
    "passcode", client.Password,
  }
 
  conn, err := stompngo.Connect(client.NetConnection, headers)
  if err != nil {
    fmt.Println(err)
  }
  
  client.StompConnection = conn
  return conn
}

func (client *Client) Connect() (conn *stompngo.Connection) {
  client.setOpts()
  client.netConnection()
  conn = client.stompConnection()
  return
}

func (client *Client) Disconnect() {
  client.StompConnection.Disconnect(stompngo.Headers{})
  client.NetConnection.Close()
}

var client Client
type Message interface {}
*/
func main() {
  conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", "61613"));
  if err != nil {
    fmt.Println(err);
    return;
  }
  defer func(){
    err = conn.Close();
    if err != nil{
      fmt.Println(err);
    }
  }();

  head := stompngo.Headers{"accept-version","1.2","host","localhost","login","guest","passcode","guest"};
  cnn, e := stompngo.Connect(conn, head);
  if e != nil{
    fmt.Println(err);
  }

  // Send message
  headers := stompngo.Headers{"destination", "/queue/bonly", "id", stompngo.Uuid()};
  err = cnn.Send(headers, "bbbbbb messssss");;
  if err != nil {
    fmt.Println(err);
  }

}