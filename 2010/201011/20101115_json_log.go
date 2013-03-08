package main
import (
  "encoding/json"
  "fmt"
)

type Glog struct{
	Ip string `json:ip`;
	Port string `json:port`;
	Realm string `json:realm`;
	Cmd string `json:cmd`;
};

func main(){
	var s Glog;
	str := `{"player_id":"51136","log_time":"20130226T094317","realm":"2","ip":"119.4.249.150","port":"14535","cmd":"67","sub_cmd_id":"0"}`;
	json.Unmarshal([]byte(str), &s);
	fmt.Println(s);
}
	