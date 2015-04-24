package main

import (
    "fmt"
    curl "github.com/andelf/go-curl"
    "flag"
    "encoding/json"
)

var url *string = flag.String("h", "http://127.0.0.1:9001", "Host url.");
var times *int = flag.Int("t", 1, "run times.");
var cmd *int = flag.Int("c", 1, "cmd.");

func main() {
  flag.Parse();
  
  for i:=0; i<*times; i++{
    switch(*cmd){
      case 1:
         test_login();
         break;
      case 2:
         test_get_realm();
         break;
      case 3:
         test_get_user_id();
         break;
      default:
         break;
    }
  }
}

// make a callback function
func callback(buf []byte, userdata interface{}) bool {
    println("DEBUG: size=>", len(buf));
    println("DEBUG: content=>", string(buf));
    
//    var ack sLoginAck;
//    json.Unmarshal([]byte(string(buf)), &ack);

    var val interface{};
    err := json.Unmarshal([]byte(string(buf)), &val);
    if err != nil{
      panic(err.Error());
    }
    
    md := val.(map[string]interface{});
    
    for k,v := range md{
      switch vv := v.(type){
        case string:
          fmt.Println(k, "[string]", vv);
        case int:
          fmt.Println(k, "[int]", vv);
        case float64:
          fmt.Println(k, "[float65]", vv);
        case []interface{}:
          fmt.Println(k, "[array]");
          for i, u := range vv{
            fmt.Println(i, u);
          }
        default:
          fmt.Println(k, "[unknow]");
      }
    }

    return true;
}

type sLogin struct{
  Cmd int `json:"cmd"`;
  User_name string `json:"user_name"`;
  Passwd string `json:"passwd"`;
  Ver int `json:"ver"`;
};
  
type sLoginAck struct{
  Cmd int `json:"cmd"`;
  Ret int `json:"ret"`;
  Update int `json:"update"`;
  Acc_id int `json:"acc_id"`;
};
  
func test_login(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  
  var mylogin sLogin;
  mylogin.Cmd = 1;
  mylogin.User_name = "test1";
  mylogin.Passwd = "passwd";
  mylogin.Ver = 1;
  
  body, err := json.Marshal(mylogin);
  if err != nil{
    panic(err.Error());
  }
  
  easy.Setopt(curl.OPT_POSTFIELDS, string(body));

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }  
}

func test_get_realm(){
//curl -d '{"cmd":3, "machine":1000, "ver":1, "acc_id":1001}' "http://127.0.0.1:9001"
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  easy.Setopt(curl.OPT_POSTFIELDS, "{\"cmd\":3, \"machine\": 1000, \"ver\": 1, \"acc_id\": 1001}");

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }  
}

func test_get_user_id(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  easy.Setopt(curl.OPT_POSTFIELDS, "{\"cmd\": 5, \"realm_id\":0, \"acc_id\": 1}");

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }  
}

//curl -d "this is body" -u "user:pass" "http://localhost/?ss=ss&qq=11"
//curl -d '{"user_id":33}' "http://127.0.0.1/8000'
