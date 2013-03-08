package main;

import (
  "flag"
  "log"
  "os"
  "fmt"
  "encoding/json"
);

/**
 全局变量 
 */
config_file string;

type DBServer struct{
	IP   string;
	Port string;
	User string;
	Passwd string;
	DbName string;
};

type Server struct{
	Port string;
	DbSrv  []DBServer;
};

/**
 生成配置模板
 */
func gen_config(){
	  var s Server
    s.DbSrv = append(s.DbSrv, DBServer{IP: "117.135.154.58", Port: "3306", User: "mysql", Passwd:"", DbName: "Paladin"})
    s.DbSrv = append(s.DbSrv, DBServer{IP: "117.135.154.26", Port: "3306", User: "mysql", Passwd:"", DbName: "Paladin"})
    b, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    }
    fmt.Println(string(b))
}

/**
 load 配置
 */
func load_config(filename string){
	fh, err := os.Open(filename);
}

 
/**
 解释参数
 */
func argv_parse(){
	help := flag.Bool("h", false, "help");
	config_gen := flag.Bool("g", false, "gen config template");
	config_file = flag.String("c", "watch.conf", "config file");
	
	flag.Parse();
	if *help == true {
		var Usage = func() {
	    fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0]);
	    flag.PrintDefaults();
	  }
	  Usage();
	}
	if *config_gen == true{
		gen_config();
	}
}

func main() {
  argv_parse();
  log.Println("server begin...");
}
	

