package main

import (
  // "encoding/json"
  "fmt"
  "github.com/pkmngo-odi/pogo/api"
  "github.com/pkmngo-odi/pogo/auth"
  "flag"
  "sort"
)

var Acc *string = flag.String("a", "h.bonly@gmail.com", "Account of google game");
var Passwd *string = flag.String("p", "", "Password");
var Type *string = flag.String("t", "google", "Account type: google/ptc");
var Auth *bool = flag.Bool("u", false, "Get Auth");

type IV struct{
  Name string;
  CP int32;
  HP int32;
  Attack int32;
  Defense int32;
  Stamina int32;
  Nickname string;
  Prefect int32;
  Creation_time_ms uint64;
};

type LstIV []IV;
func (ls LstIV) Len() int{
  return len(ls);
}
func (ls LstIV) Swap(i, j int){
  ls[i], ls[j] = ls[j], ls[i];
}
func (ls LstIV) Less(i, j int) bool{
  return ls[i].Creation_time_ms < ls[j].Creation_time_ms;
}

var lstIv LstIV;


func main() {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println(err);
    }
  }();

  flag.Parse();

  if len(*Acc)<=0 || len(*Passwd)<=0{
    return;
  }

  if *Auth {
    Get_Auth();
    return;
  }  

  // Initialize a new authentication provider to retrieve an access token
  // provider, err := auth.NewProvider("ptc", "h.bonly@gmail.com", "hay111")
  provider, err := auth.NewProvider(*Type, *Acc, *Passwd)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Set the coordinates from where you're connecting
  location := &api.Location{
    Lon: 37.79633928485233,
    Lat: -122.40627765655516,
    Alt: 0.0,
  }

  feed := api.NewFeed(&api.VoidReporter{});

  // Start new session and connect
  session := api.NewSession(provider, location, feed, false)
  session.Init()

  // Start querying the API
  player, err := session.GetPlayer()
  if err != nil {
    fmt.Println(err)
    return
  }

  // out, err := json.Marshal(player)
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }

  if player.Success == false{
    fmt.Println("玩家查询失败");
    return;
  }
  // fmt.Println("玩家：", string(out));
  fmt.Println("姓名：", player.PlayerData.Username);
  fmt.Println("队伍：", player.PlayerData.Team);
 

  stuff, err := session.GetInventory();
  if err != nil{
  	fmt.Println(err);
  	return;
  }

  // stu, err := json.Marshal(stuff); //结构转array
  // if err != nil{
  // 	fmt.Println(err);
  // 	return;
  // } 
  // fmt.Println("物品：", string(stu));

  lstIv = make([]IV,0);

  for _, it := range stuff.InventoryDelta.InventoryItems{
    // fmt.Printf("[%d]: %+v\n", idx, it.InventoryItemData.PokemonData);
    if it.InventoryItemData.PokemonData != nil{
      if it.InventoryItemData.PokemonData.Cp==0{
        // fmt.Printf("[%d]:%15v\n", idx, "蛋");
      }else{
        prefect :=Get_perc(it.InventoryItemData.PokemonData.IndividualAttack,
            it.InventoryItemData.PokemonData.IndividualDefense,
            it.InventoryItemData.PokemonData.IndividualStamina);

        lstIv = append(lstIv, IV{
            it.InventoryItemData.PokemonData.PokemonId.String(),
            it.InventoryItemData.PokemonData.Cp,
            it.InventoryItemData.PokemonData.StaminaMax,
            it.InventoryItemData.PokemonData.IndividualAttack,
            it.InventoryItemData.PokemonData.IndividualDefense,
            it.InventoryItemData.PokemonData.IndividualStamina,
            it.InventoryItemData.PokemonData.Nickname,
            prefect,
            it.InventoryItemData.PokemonData.CreationTimeMs,
          });
        /*
        fmt.Printf("[%3d]:%15v  CP:%5v  HP:%5v  Attack:%5v  Defense:%5v  Stamina:%5v  Name:%15v\tPre:%3v%%\n", idx, 
          it.InventoryItemData.PokemonData.PokemonId,
          // it.InventoryItemData.PokemonData.CpMultiplier,
          it.InventoryItemData.PokemonData.Cp,
          it.InventoryItemData.PokemonData.StaminaMax,
          it.InventoryItemData.PokemonData.IndividualAttack,
          it.InventoryItemData.PokemonData.IndividualDefense,
          it.InventoryItemData.PokemonData.IndividualStamina,
          it.InventoryItemData.PokemonData.Nickname,
          prefect));
        */
      }
    }
  }
  
  sort.Sort(lstIv);
  for idx, it := range lstIv{
        fmt.Printf("[%3d]:%15v  HP:%5v  Attack:%5v  Defense:%5v  Stamina:%5v  CP:%5v  Name:%15v\tPre:%3v%%\n", idx, 
          it.Name,
          it.HP,
          it.Attack,
          it.Defense,
          it.Stamina,
          it.CP,
          it.Nickname,
          it.Prefect);    
  }

}

func Get_perc(atk int32, def int32, sta int32) (pre int32){
    return (int32) (atk + def + sta) * 100 / 45.0 ;
}

func Get_Auth(){
  // Initialize a new authentication provider to retrieve an access token
  provider, err := auth.NewProvider(*Type, *Acc, *Passwd)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Set the coordinates from where you're connecting
  location := &api.Location{
    Lon: 0.0,
    Lat: 0.0,
    Alt: 0.0,
  }

  feed := api.NewFeed(&api.VoidReporter{});

  // Start new session and connect
  session := api.NewSession(provider, location, feed, false)
  session.Init()

  token, err := provider.Login();
  if err != nil{
    fmt.Println(err);
    return;
  }
  fmt.Println(token);  
}
//GOOS=darwin GOARCH=amd64
//CGO_ENABLED=0 GOOS=windows GOARCH=amd64