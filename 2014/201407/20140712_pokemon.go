package main

import (
  "encoding/json"
  "fmt"
  "github.com/pkmngo-odi/pogo/api"
  "github.com/pkmngo-odi/pogo/auth"
)

func main() {

  // Initialize a new authentication provider to retrieve an access token
  // provider, err := auth.NewProvider("ptc", "h.bonly@gmail.com", "hay111")
  provider, err := auth.NewProvider("google", "h.bonly@gmail.com", "hay111")
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

  // Start new session and connect
  session := api.NewSession(provider, location, false)
  session.Init()

  // Start querying the API
  player, err := session.GetPlayer()
  if err != nil {
    fmt.Println(err)
    return
  }

  out, err := json.Marshal(player)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(string(out))

  stuff, err := session.GetInventory();
  if err != nil{
  	fmt.Println(err);
  	return;
  }

  stu, err := json.Marshal(stuff);
  if err != nil{
  	fmt.Println(err);
  	return;
  } 
  fmt.Println(string(stu));
}