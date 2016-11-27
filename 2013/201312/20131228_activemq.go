package main

import (
  "os"
  "fmt"
  // "errors"
  "net"
  "encoding/json"
  "github.com/gmallard/stompngo"
)

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

func main() {
  conn := client.Connect()
  defer client.Disconnect()

  // Subscribe to a channel
  headers := stompngo.Headers{"destination", "/topic/TRAIN_MVT_ALL_TOC", "id", client.Uuid}
  channel, err := conn.Subscribe(headers)
  if err != nil {
    fmt.Println(err)
  }

  for {
    data := <-channel
    if data.Error != nil {
      fmt.Println(data.Error)
    }
    
    var messages []Message

    err := json.Unmarshal(data.Message.Body, &messages)
    if err != nil {
      fmt.Println(err)
    }

    for _, message := range messages {
      fmt.Println(message)
    }

  }
}