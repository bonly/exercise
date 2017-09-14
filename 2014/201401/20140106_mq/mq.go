package main

import (
"os"
"fmt"
"net"
"github.com/gmallard/stompngo"
"encoding/json"
"time"
)

type CLog struct{
	ConfigItemCode string `json:"configItemCode"`;
	LogType string `json:"logType"`;
	LogLevel string `json:"logLevel"`;
	LogContent string `json:"logContent"`;
	LogRemark string `json:"logRemark"`;
	LogCode string `json:"logCode"`;
	CreateTime int64 `json:"createTime"`;
};

type Client struct{
	Host string;
	Port string;
	User string;
	Password string;
	Uuid string;
	Net  net.Conn;
	SNet *stompngo.Connection;
	QName string;
};

func (client *Client) set_opts(){
  client.Host = "localhost";
  client.Port = "61613";
  client.QName = "q.notify";
  client.User = "";
  client.Password = "";
  client.Uuid = stompngo.Uuid();
  
  cli_host := os.Getenv("MQ_HOST")
  if cli_host != "" {
    client.Host = cli_host;
  }

  cli_port := os.Getenv("MQ_PORT");
  if cli_port != "" {
    client.Port = cli_port;
  }

  cli_user := os.Getenv("MQ_USER");
  if cli_user != "" {
    client.User = cli_user;
  }

  cli_password := os.Getenv("MQ_PASSWORD");
  if cli_password != "" {
    client.Password = cli_password;
  }

  cli_uuid := os.Getenv("MQ_UUID");
  if cli_uuid != "" {
    client.Uuid = cli_uuid;
  }	

  cli_qname := os.Getenv("MQ_NAME");
  if cli_qname != "" {
    client.QName = cli_qname;
  }	  
}

func (client *Client) connect() (conn net.Conn, err error) {
  conn, err = net.Dial("tcp", net.JoinHostPort(client.Host, client.Port));
  if err != nil {
    fmt.Println(err);
    return nil, err;
  }

  client.Net = conn;
  return;
}

func (client *Client) sconnect() (*stompngo.Connection) {
  headers := stompngo.Headers{
    "accept-version", "1.2", 
    "host", client.Host,
    "login", client.User,
    "passcode", client.Password,
  };
 
  conn, err := stompngo.Connect(client.Net, headers);
  if err != nil {
    fmt.Println(err);
  }
  
  client.SNet = conn;
  return conn;
}

func (client *Client) Connect() (conn *stompngo.Connection) {
  client.set_opts();
  client.connect();
  conn = client.sconnect();
  return;
}

func (client *Client) Disconnect() {
  client.SNet.Disconnect(stompngo.Headers{});
  client.Net.Close();
}

var client Client;
type Message interface{};

//export Msg
func Msg(configItemCode string, logType string, logLevel string, logContent string,
		logRemark string, logCode string){
	conn := client.Connect();
	defer client.Disconnect();

	headers := stompngo.Headers{"destination", client.QName, "id", client.Uuid};

	var clog CLog;
	clog.CreateTime = time.Now().Unix();
	clog.ConfigItemCode = configItemCode;
	clog.LogType = logType;
	clog.LogLevel = logLevel;
	clog.LogContent = logContent;
	clog.LogRemark = logRemark;
	clog.LogCode = logCode;

	rs, err := json.Marshal(clog);
	if err != nil{
		fmt.Println(err);
		return;
	}
	conn.Send(headers, string(rs));
}

func main(){
	Msg("item", "error", "44", "出错了", "remark", "12234");
}