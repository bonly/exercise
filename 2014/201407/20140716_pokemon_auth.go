package main

import (
  // "encoding/json"
  "fmt"
  "github.com/pkmngo-odi/pogo/api"
  "github.com/pkmngo-odi/pogo/auth"
)

func main() {

  // Initialize a new authentication provider to retrieve an access token
  provider, err := auth.NewProvider("google", "fallingwood07@gmail.com", "qingfeng")
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

  // Start new session and connect
  session := api.NewSession(provider, location, false)
  session.Init()

  token, err := provider.Login();
  if err != nil{
    fmt.Println(err);
    return;
  }
  fmt.Println(token);
}