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
var player_id *int = flag.Int("p", 1, "player id.");

func main() {
  flag.Parse();
  
  for i:=0; i<*times; i++{
    switch(*cmd){
      case 1:
         test_login();
         break;
      case 3:
         test_get_realm();
         break;
      case 5:
         test_get_user_id();
         break;
      case 7:
         test_apply_account();
         break;
      case 9:
         test_get_init_mech();
         break;   
      case 11:
         test_apply_player();
         break;     
      case 13:
         test_sync_cust_user();
         break;
      case 15:
         test_get_player_info();
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
    
    var val interface{};
    err := json.Unmarshal([]byte(string(buf)), &val);
    if err != nil{
      panic(err.Error());
    }
    
    md := val.(map[string]interface{});
    
    for k,v := range md{
      switch vv := v.(type){
        case string:
          fmt.Println(k, "[string]:", vv);
        case int:
          fmt.Println(k, "[int]:", vv);
        case float64:
          fmt.Println(k, "[float65]:", vv);
        case []interface{}:
          fmt.Println(k, "[array]:");
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
   
func test_login(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  
  var mylogin sLogin;
  mylogin.Cmd = 1;
  mylogin.User_name = "snowlion";
  mylogin.Passwd = "snowlion";
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

type sGet_realm struct{
  Cmd int `json:"cmd"`;
  Machine int `json:"machine"`;
  Ver int `json:"ver"`;
  Acc_id int `json:"acc_id"`;
};

func test_get_realm(){
//curl -d '{"cmd":3, "machine":1000, "ver":1, "acc_id":1001}' "http://127.0.0.1:9001"
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);

  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  var getRealm sGet_realm;
  getRealm.Cmd=3;
  getRealm.Machine=1000;
  getRealm.Ver=1;
  getRealm.Acc_id=1001;
  
  body, err := json.Marshal(getRealm);
  if err != nil{
    panic(err.Error());
  }
  easy.Setopt(curl.OPT_POSTFIELDS,  string(body));

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }  
}

type sGet_user_id struct{
  Cmd int `json:"cmd"`;
  Realm_id int `json:"realm_id"`;
  Acc_id int `json:"acc_id"`;
};

func test_get_user_id(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  var Get_user_id sGet_user_id;
  Get_user_id.Cmd=5;
  Get_user_id.Realm_id=0;
  Get_user_id.Acc_id=1;  
  body, err := json.Marshal(Get_user_id);
  if err != nil{
    panic(err.Error());
  }
  easy.Setopt(curl.OPT_POSTFIELDS, string(body));

  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }  
}

type sApply_account struct{
  Cmd int `json:"cmd"`;
  Acc_name string `json:"acc_name"`;
  Passwd string `json:"passwd"`;
};  

func test_apply_account(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  var Apply_account sApply_account;
  Apply_account.Cmd=7;
  Apply_account.Acc_name="abc@gmail.com";
  Apply_account.Passwd="passwd";  
  body, err := json.Marshal(Apply_account);
  if err != nil{
    panic(err.Error());
  }
  easy.Setopt(curl.OPT_POSTFIELDS, string(body));

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }    
}

type sGet_init_mech struct{
  Cmd int `json:"cmd"`;
  Acc_id int `json:"acc_id"`;
};  

func test_get_init_mech(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);


  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  var InitMech sGet_init_mech;
  InitMech.Cmd=9;
  InitMech.Acc_id=1001;
  body, err := json.Marshal(InitMech);
  if err != nil{
    panic(err.Error());
  }
  easy.Setopt(curl.OPT_POSTFIELDS, string(body));

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }    
}

type sApply_player struct{
  Cmd int `json:"cmd"`;
  Acc_id int `json:"acc_id"`;
  Mech_id int `json:"mech_id"`;
  Player_name string `json:"player_name"`;
  Realm_id int `json:"realm_id"`;  
};  

func test_apply_player(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);

  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  var player sApply_player;
  player.Cmd=11;
  player.Acc_id=1001;
  player.Mech_id=11000;
  player.Player_name="无敌";
  player.Realm_id=0;
  body, err := json.Marshal(player);
  if err != nil{
    panic(err.Error());
  }
  easy.Setopt(curl.OPT_POSTFIELDS, string(body));

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }    
}

type sSyncCustUser struct{
  Cmd int `json:"cmd"`;
  Acc_id int `json:"acc_id"`;
  User_id int `json:"user_id"`;
  Realm_id int `json:"realm_id"`;  
};  

func test_sync_cust_user(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);

  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  var player sSyncCustUser;
  player.Cmd=13;
  player.Acc_id=1;
  player.User_id=1000;
  player.Realm_id=0;
  body, err := json.Marshal(player);
  if err != nil{
    panic(err.Error());
  }
  easy.Setopt(curl.OPT_POSTFIELDS, string(body));

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }    
}

type sQryPlayer struct{
  Cmd int `json:"cmd"`;
  Player_id int `json:"player_id"`;
};

func test_get_player_info(){
  easy := curl.EasyInit();
  defer easy.Cleanup();

  fmt.Println(*url);
  easy.Setopt(curl.OPT_URL, *url);

  easy.Setopt(curl.OPT_WRITEFUNCTION, callback);
  var player sQryPlayer;
  player.Cmd=15;
  player.Player_id=*player_id;

  body, err := json.Marshal(player);
  if err != nil{
    panic(err.Error());
  }
  easy.Setopt(curl.OPT_POSTFIELDS, string(body));

  
  if err := easy.Perform(); err != nil {
      fmt.Printf("ERROR: %v\n", err);
  }      
}

//curl -d "this is body" -u "user:pass" "http://localhost/?ss=ss&qq=11"
//curl -d '{"user_id":33}' "http://127.0.0.1/8000'
