/*
auth: bonly
create: 2016.9.27
desc: 门锁主动的操作，回调给OMS
*/
package callback

import(
"proto"
"oms"
"config"
"net/http"
"fmt"
"io/ioutil"
"strings"
log "glog"
"time"
"bcd"
)

func Run()(cbc chan interface{}){
	cbc = make(chan interface{});
	go func(){
		for config.Run {
			select {
				case pack, ok := <- cbc :{
					if !ok {
						return;
					}
					log.Infof("收到门锁主动数据\n");
					log_data := pack.(*proto.Lock_upload_log);
					go func(){
						process(log_data);
					}();
				}
			}
		}
	}();
	return cbc;
}

func process(data *proto.Lock_upload_log){
	var tan oms.Lock_log_REQ;
	tan.Cmd = *data;

	cmd := (uint8)(tan.Cmd.Data.Open_type);
	switch {
		case cmd  == proto.OPEN_DOOR_WITH_PASSWD:{
			log.Infof("门锁上传密码开门记录\n");

			loc, _ := time.LoadLocation("Local");
			tn := time.Date( (int)(tan.Cmd.Data.DateTime[0])+2000, //年
				             (time.Month)(tan.Cmd.Data.DateTime[1]), //月
				             (int)(tan.Cmd.Data.DateTime[2]), //日
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[3]))), //时
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[4]))), //分
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[5]))), //秒
				             0, loc);
			var passwd string;
			for idx := 0; idx < (int)(tan.Cmd.Data.SN_len); idx++{
				char := fmt.Sprintf("%d", tan.Cmd.Data.SN[idx]);
				passwd += char;
			}
			cbc := &oms.Lock_log_REQ{
				Name : "PasswordOpenRoom_REQ",
				LockId : fmt.Sprintf("%d", tan.Cmd.Command.Head().LockID),
				MidId : tan.Cmd.MidId,
				Param : passwd,
				DateTime : tn.Format("2006-01-02 15:04:05"),
			};
			pack, _ := cbc.Encode();
			Post(*config.Callback, string(pack));
			break;
		}
		case cmd == proto.DOOR_LOW_POWER:{
			log.Infof("门锁上传低电压警报\n");

			loc, _ := time.LoadLocation("Local");
			tn := time.Date( (int)(tan.Cmd.Data.DateTime[0])+2000, //年
				             (time.Month)(tan.Cmd.Data.DateTime[1]), //月
				             (int)(tan.Cmd.Data.DateTime[2]), //日
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[3]))), //时
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[4]))), //分
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[5]))), //秒
				             0, loc);
			var passwd string;
			for idx := 0; idx < (int)(tan.Cmd.Data.SN_len); idx++{
				char := fmt.Sprintf("%d", tan.Cmd.Data.SN[idx]);
				passwd += char;
			}
			cbc := &oms.Lock_log_REQ{
				Name : "Alert_REQ",
				LockId : fmt.Sprintf("%d", tan.Cmd.Command.Head().LockID),
				MidId : tan.Cmd.MidId,
				Param : passwd,
				DateTime : tn.Format("2006-01-02 15:04:05"),
			};
			pack, _ := cbc.Encode();
			Post(*config.Callback, string(pack));
			break;			
		}
		case cmd == proto.FORBIT_OPEN_DOOR:{
			log.Infof("门锁上传被暴力破解\n");

			loc, _ := time.LoadLocation("Local");
			tn := time.Date( (int)(tan.Cmd.Data.DateTime[0])+2000, //年
				             (time.Month)(tan.Cmd.Data.DateTime[1]), //月
				             (int)(tan.Cmd.Data.DateTime[2]), //日
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[3]))), //时
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[4]))), //分
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[5]))), //秒
				             0, loc);
			var passwd string;
			for idx := 0; idx < (int)(tan.Cmd.Data.SN_len); idx++{
				char := fmt.Sprintf("%d", tan.Cmd.Data.SN[idx]);
				passwd += char;
			}
			cbc := &oms.Lock_log_REQ{
				Name : "BruteForce_REQ",
				LockId : fmt.Sprintf("%d", tan.Cmd.Command.Head().LockID),
				MidId : tan.Cmd.MidId,
				Param : passwd,
				DateTime : tn.Format("2006-01-02 15:04:05"),
			};
			pack, _ := cbc.Encode();
			Post(*config.Callback, string(pack));
			break;			
		}		
		case cmd == proto.MULTI_OPEN_DOOR_ERR:{
			log.Infof("门锁被尝试密码开门\n");

			loc, _ := time.LoadLocation("Local");
			tn := time.Date( (int)(tan.Cmd.Data.DateTime[0])+2000, //年
				             (time.Month)(tan.Cmd.Data.DateTime[1]), //月
				             (int)(tan.Cmd.Data.DateTime[2]), //日
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[3]))), //时
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[4]))), //分
				             (int)(bcd.BcdToInt((int)(tan.Cmd.Data.DateTime[5]))), //秒
				             0, loc);
			var passwd string;
			for idx := 0; idx < (int)(tan.Cmd.Data.SN_len); idx++{
				char := fmt.Sprintf("%d", tan.Cmd.Data.SN[idx]);
				passwd += char;
			}
			cbc := &oms.Lock_log_REQ{
				Name : "OpenRoomPassWordErr_REQ",
				LockId : fmt.Sprintf("%d", tan.Cmd.Command.Head().LockID),
				MidId : tan.Cmd.MidId,
				Param : passwd,
				DateTime : tn.Format("2006-01-02 15:04:05"),
			};
			pack, _ := cbc.Encode();
			Post(*config.Callback, string(pack));
			break;			
		}				
		default :{
			log.Warningf("未知的记录类型[%X]\n", cmd);
			break;
		}
	}
}

//post oms
func Post(srv string, data string){
	log.V(99).Infof("Send oms: %s\n", data);
	request, err := http.NewRequest("POST",
		srv, strings.NewReader(data));
	if err != nil{
		log.Warningf("REQ OMS %s\n", err.Error());
		return;
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	request.Header.Set("Content-Type", "application/json");
	request.Header.Set("Accept", "application/json");

	cli := &http.Client{};
	resp, err := cli.Do(request);
	if err != nil{
		log.Warningf("POST %s\n", err.Error());
		return;
	}	
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		log.Warningf("body %s\n", err.Error());
		return;
	}
	log.Infof("Recv oms: %s\n", string(body));
}