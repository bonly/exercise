package main
 
import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "io/ioutil"
    "bytes"
    "github.com/bitly/go-nsq"
)
 
type handle struct {
    host string;
    port string;
};

type myReader struct {
    *bytes.Buffer;
};

func (m myReader) Close() error { return nil };

type myNsq struct{
    NsqCfg *nsq.Config; 
    NsqPro *nsq.Producer;
};

var MyNsq myNsq;

func (this *myNsq)Nsq() *nsq.Producer{
    if this.NsqPro == nil{
        this.NsqCfg = nsq.NewConfig();
        this.NsqPro, _ = nsq.NewProducer("120.25.106.243:4150", this.NsqCfg);
    }
    return this.NsqPro;
}

func (this *myNsq)Stop(){
    if this.NsqPro != nil{
        this.NsqPro.Stop();
    }
}

func (this *myNsq)Distribute_msg(msg []byte){
  w := MyNsq.Nsq();

  err := w.Publish("write_test", msg);
  if err != nil {
      log.Panic("Could not connect");
  }
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) { //服务处理响应
    remote, err := url.Parse("http://" + this.host + ":" + this.port); //解释地址
    if err != nil {
        panic(err);
    }
    
    buf, _ := ioutil.ReadAll(r.Body)
    rdr1 := myReader{bytes.NewBuffer(buf)}
    rdr2 := myReader{bytes.NewBuffer(buf)}

    go process(w, rdr1);
    r.Body = rdr2;
    proxy := httputil.NewSingleHostReverseProxy(remote); //新建一个反向服务请求
    proxy.ServeHTTP(w, r); //调服务的请求响应
}
 
func startServer() {
    defer MyNsq.Stop();

    //被代理的服务器host和port
    h := &handle{host: "120.25.106.243", port: "7008"}; //建一个服务
    err := http.ListenAndServe(":7009", h); //侦听
    if err != nil {
        log.Fatalln("ListenAndServe: ", err);
    }
}

func process(wr http.ResponseWriter, rd myReader){
    body, err := ioutil.ReadAll(rd);
    if err != nil{
        log.Println(err);
        return;
    }
    log.Println(string(body));
    MyNsq.Distribute_msg(body);
    return;
} 

func main() {
    startServer();
}

