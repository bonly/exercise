package main
 
import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "io/ioutil"
    "bytes"
)
 
type handle struct {
    host string
    port string
};

type myReader struct {
    *bytes.Buffer
}

func (m myReader) Close() error { return nil };

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) { //服务处理响应
    remote, err := url.Parse("http://" + this.host + ":" + this.port); //解释地址
    if err != nil {
        panic(err);
    }
    
    buf, _ := ioutil.ReadAll(r.Body)
    rdr1 := myReader{bytes.NewBuffer(buf)}
    rdr2 := myReader{bytes.NewBuffer(buf)}

    go print(w, rdr1);
    r.Body = rdr2;
    proxy := httputil.NewSingleHostReverseProxy(remote); //新建一个反向服务请求
    proxy.ServeHTTP(w, r); //调服务的请求响应
}
 
func startServer() {
    //被代理的服务器host和port
    h := &handle{host: "120.25.106.243", port: "7008"}; //建一个服务
    err := http.ListenAndServe(":7009", h); //侦听
    if err != nil {
        log.Fatalln("ListenAndServe: ", err);
    }
}

func print(wr http.ResponseWriter, rd myReader){
    body, err := ioutil.ReadAll(rd);
    if err != nil{
        log.Println(err);
        return;
    }
    log.Println(string(body));
    return;
} 

func main() {
    startServer();
}



/*
type myReader struct {
    *bytes.Buffer
}

// So that it implements the io.ReadCloser interface
func (m myReader) Close() error { return nil } 

buf, _ := ioutil.ReadAll(r.Body)
rdr1 := myReader{bytes.NewBuffer(buf)}
rdr2 := myReader{bytes.NewBuffer(buf)}

doStuff(rdr1)
r.Body = rdr2 // OK since rdr2 implements the io.ReadCloser interface

// Now the program can continue oblivious to the fact that
// r.Body was ever touched.

http://stackoverflow.com/questions/23070876/reading-body-of-http-request-without-modifying-request-state
*/