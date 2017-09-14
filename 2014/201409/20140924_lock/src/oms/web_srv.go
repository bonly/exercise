package oms
import (
"fmt"
"flag"
"net/http"
"log"
"io/ioutil"
"sync"
"config"
)

type SRV struct{
};

func (this *SRV)Srv(){
	flag.Parse();

	http.HandleFunc("/", this.home);
	err := http.ListenAndServe(*config.Oms_srv, nil)
	if err != nil {
		log.Fatal(err);
	}
}

func (this *SRV)home(wr http.ResponseWriter, rq *http.Request){
	wr.Header().Set("Access-Control-Allow-Origin", "*");
	wr.Header().Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	
	body, err := ioutil.ReadAll(rq.Body);
	if err != nil{
		fmt.Printf("%v\n", err);
		return;
	}

	if len(body) <= 0{
		fmt.Printf("消息体空\n");
		return;
	}

	var wg sync.WaitGroup;
	wg.Add(2);
	ret := this.process(body, &wg);
	defer func(){
		close(ret);
	}();
	// wr.Write([]byte("OK"));
	go func(){
		defer wg.Done();
		for {
			select{
				case result := <- ret:{
					if result == nil{
						fmt.Printf("收到nil\n");
						return;
					}
					//todo 回写到客户端
					fmt.Printf("%+v\n", result);
					wr.Write([]byte("OK"));
					return;
				}
			}
		}
	}();
	wg.Wait();
	fmt.Printf("结束处理\n");
}

func (this *SRV)process(pack []byte, wg *sync.WaitGroup)(ret chan interface{}){
	ret = make(chan interface{});
	go func(){
		defer wg.Done();
		var cmd PMS_CMD;
		cmd = &PMS_Command{};
		err := cmd.Decode(pack);
		if err != nil{
			fmt.Printf("xml wrong %v\n", err);
			return; //todo 返回错误到chan
		}
		cmd_name := *(cmd.(*PMS_Command).Name);
		fmt.Printf("命令: %s\n", cmd_name);

		switch(cmd_name){
			case "Add_OpenUser_REQ":{
				fmt.Printf("新增用户\n");
				// cmd = &Add_OpenUser_REQ{};
				break;
			}
			case "Delete_OpenUser_REQ":{
				fmt.Printf("删除用户\n");
				break;
			}
			case "OpenLock_REQ":{
				fmt.Printf("远程开门\n");
				cmd = &OpenLock_REQ{};
				break;
			}			
		}
		err = cmd.Decode(pack);
		if err != nil{
			return; //todo 返回错误到chan
		}
		cmd.Process(ret);
	}();
	return ret;
}