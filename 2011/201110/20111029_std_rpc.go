package main
 
import (
        "net/rpc"
        "net/http"
        "log"
        "net"
        "time"
)
 
 
type Args struct {
        A, B int
}
 
type Arith int;  //作为被调用类定义，必须为int
 
 //首字母必须大写，必须有两个参数，第一个是接收参数，第二个是返回参数，error返回值
func (t *Arith) Multiply(args *Args, reply *([]string)) error {
        *reply = append(*reply, "test")
        return nil
}
 
func main() {
        arith := new(Arith)
 
        rpc.Register(arith); // 注册服务
        rpc.HandleHTTP(); // 注册协议
         
        l, e := net.Listen("tcp", ":1234")
        if e != nil {
                log.Fatal("listen error:", e)
        }
        go http.Serve(l, nil)
        time.Sleep(5 * time.Second)
 
        client, err := rpc.DialHTTP("tcp", "127.0.0.1" + ":1234")
        if err != nil {
                log.Fatal("dialing:", err)
        }
         
        args := &Args{7,8}; //调用参数，其实没做什么用
        reply := make([]string, 10)
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        log.Println(reply)
}

/*
rpc的标准调用方式
*/
