//+build ignore

package main

import (
  "fmt"
  "log"
  "net"
  "net/rpc"
  "net/rpc/jsonrpc"
  )

type Args struct {
        A, B int
}

type Reply struct {
        C int
}

type Arith int

func (t *Arith) Mul(args *Args, reply *Reply) error {
        reply.C = args.A * args.B
        return nil
}

func startServer(){
    arith := new(Arith);

    rpc.Register(arith);

    //rpc.HandleHTTP (rpc.DefaultRPCPath, rpc.DefaultDebugPath);

    l, e := net.Listen("tcp", ":8222");
    if e != nil {
      log.Fatal("listen error: ", e);
    }

    for {
      conn, err := l.Accept();
      if err != nil {
        log.Fatal(err);
      }

      go rpc.ServeCodec(jsonrpc.NewServerCodec(conn));
    }
}

func main(){
    go startServer(); //服务端

    cli, err := jsonrpc.Dial("tcp", "localhost:8222"); //使用jsonrpc中的

    if err != nil {
       panic(err);
    }
    defer cli.Close();

    args := Args{7, 8};
    
    var reply Reply;

    for i:=0; i<1; i++{
      err = cli.Call("Arith.Mul", args, &reply);
      if err != nil{
         log.Fatal("arith error: ", err);
      }
      fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply);
    }
}
