package log_pkg
import (
  //"fmt"
  "labix.org/v2/mgo"
  //"labix.org/v2/mgo/bson"
  "encoding/json"
)

type Glog struct{
	Ip string `json:ip`;
	Port string `json:port`;
	Realm string `json:realm`;
	Cmd string `json:cmd`;
};

type MyDb struct{
	Session *mgo.Session;
};

func (db *MyDb)Open(str string){
	session, err := mgo.Dial(str);
	if err != nil {
		panic(err);
	}
  session.SetMode(mgo.Monotonic, true); //可选的模式切换
  db.Session = session;
}

func (db *MyDb)Close(){
	db.Session.Close();
}

func (db *MyDb)Write_log(data []byte){	
	c := db.Session.DB("test").C("gsrv");
	var log_data Glog;
	json.Unmarshal(data, &log_data);
	err := c.Insert(&log_data);
	if err != nil {
		panic(err);
	}
}
