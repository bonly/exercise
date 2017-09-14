package main

import (
  "fmt"
  p "github.com/pkmngo-odi/pogo-protos"
)

func main() {
  fmt.Println(p.TeamColor_BLUE);
  fmt.Println(p.TeamColor_RED);
}

/*
$ git clone git@github.com:AeonLucid/POGOProtos.git
$ cd POGOProtos
$ python ./compile_single.py -l=go --out_path=$GOPATH/src/github.com/pkmngo-odi/pogo-protos --go_root_package=github.com/pkmngo-odi/pogo-protos
*/