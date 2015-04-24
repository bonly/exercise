package main

import (
  "os"
  // "fmt"
  "log"
  "sync"
  "syscall"
  "os/signal"
  "io/ioutil"
  "encoding/json"
  "time"
  // "github.com/donnie4w/go-logger/logger"
)

/// 配置文件内容
type Config struct{
	Srv string;
	DB  string;
	Log_path string;
	Log_file string;
	Stdout bool;
}


var config *Config;
var configLock = new(sync.RWMutex);
var config_file *string;

/// 加载配置
func loadConfig(fail bool){
	file, err := ioutil.ReadFile(*config_file);
	if err != nil{
		log.Println("open config: ", err);
		//logger.Fatal(fmt.Sprintf("open config: %s", err));
		if fail{
			os.Exit(1);
		}
	}

	temp := new(Config);
	if err = json.Unmarshal(file, temp); err != nil {
		log.Println("parse config: ", err);
		//logger.Fatal(fmt.Sprintf("parse config: %s", err));
		if fail {
			os.Exit(1);
		}
	}

	configLock.Lock();
	config = temp;
	configLock.Unlock();
}

/// 访问单体实例
func GetConfig() *Config{
	configLock.RLock();
	defer configLock.RUnlock();
	return config;
}

/// 安装初始
func Init(file_name *string){
	if file_name == nil{
		tmpfile := "config.json";
		file_name = &tmpfile;
	}

	config_file = file_name;

	loadConfig(true);
	s := make(chan os.Signal, 1);
	signal.Notify(s, syscall.SIGUSR2);

	go func(){
		for {
			<-s;
			loadConfig(false);
			log.Println("Reloaded");
		}
	}();

}

func test_check(){
	log.Println("Log_path=", GetConfig().Log_path);
}
func main(){
	go Init(nil);
	time.Sleep(500);
	for ;;{
		time.Sleep(5000);
		test_check();
	}

}
