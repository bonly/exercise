package main

import (
    "fmt"
    curl "github.com/andelf/go-curl"
    "flag"
)

var url *string = flag.String("h", "http://127.0.0.1:9001", "Host url.");
var times *int = flag.Int("t", 1, "run times.");

func main() {
  flag.Parse();
  
  for i:=0; i<*times; i++{
    test_login();
    test_get_realm();
    test_get_user_id();
  }
}

// make a callback function
func callback(buf []byte, userdata interface{}) bool {
    println("DEBUG: size=>", len(buf));
    println("DEBUG: content=>", string(buf));
    return true;
}

func test_login(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  easy.Setopt(curl.OPT_POSTFIELDS, "{\"cmd\":1, \"acc_id\":\"abc@gmail.com\", \"passwd\": \"mypasswd\"}");

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }  
}

func test_get_realm(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  easy.Setopt(curl.OPT_POSTFIELDS, "{\"cmd\":3, \"machine\": 1000, \"ver\": 1}");

  
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
  easy.Setopt(curl.OPT_POSTFIELDS, "{\"cmd\": 5, \"realm_id\":1, \"acc_id\": \"abc@gmail.com\"}");

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }  
}

//curl -d "this is body" -u "user:pass" "http://localhost/?ss=ss&qq=11"
//curl -d '{"user_id":33}' "http://127.0.0.1/8000'