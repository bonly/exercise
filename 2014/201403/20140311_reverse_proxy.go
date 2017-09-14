package main
 
import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)
 
type handle struct {
    host string
    port string
}
 
func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) { //服务处理响应
    remote, err := url.Parse("http://" + this.host + ":" + this.port) //解释地址
    if err != nil {
        panic(err)
    }
    proxy := httputil.NewSingleHostReverseProxy(remote) //新建一个反向服务请求
    proxy.ServeHTTP(w, r) //调服务的请求响应
}
 
func startServer() {
    //被代理的服务器host和port
    h := &handle{host: "127.0.0.1", port: "80"} //建一个服务
    err := http.ListenAndServe(":8888", h) //侦听
    if err != nil {
        log.Fatalln("ListenAndServe: ", err)
    }
}
 
func main() {
    startServer()
}
