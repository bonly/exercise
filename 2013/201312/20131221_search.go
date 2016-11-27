package main 

import (
"fmt"
"encoding/json"
"log"
"sort"
)

var data = string(
`{"retCode":21020000,"msg":"操作成功","data":[{"calendarDate":"2015-11-01","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-02","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-03","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-04","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-05","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-06","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-07","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-08","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-09","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-10","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-11","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-12","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-13","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-14","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-15","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-16","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-17","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-18","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-19","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-20","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-21","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-22","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-23","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-24","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-25","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-26","price":1,"state":2,"roomId":1},{"calendarDate":"2015-11-27","price":1,"state":1,"roomId":1},{"calendarDate":"2015-11-28","price":123,"state":3,"roomId":1},{"calendarDate":"2015-11-29","price":1,"state":3,"roomId":1},{"calendarDate":"2015-11-30","price":1,"state":1,"roomId":1}]}`);




type CStat struct{
  State int `json:"state"`;
  Calendar string `json:"calendarDate"`;
  Price int `json:"price"`;
  RoomId int `json:"roomId"`;
};

type Ret struct{
	Msg string `json:"msg"`;
	RetCode int64 `json:"retCode"`;
	Data VC `json:"data"`;
};

type VC []CStat; //定义搜索用类

func (c VC) Len() int{
	return len(c);
}

func (c VC) Swap(i, j int){
	c[i], c[j] = c[j], c[i];
}

func (c VC) Less(i, j int) bool{
	return c[i].Calendar < c[j].Calendar;
}

func main(){
	var dt Ret;
	err := json.Unmarshal([]byte(data), &dt);
	if err != nil{
		log.Println(fmt.Sprint(err));
		return;
	}

	fmt.Println(sort.IsSorted(dt.Data));
	fmt.Println(dt.Data);

	pk := sort.Search(len(dt.Data), func(i int)bool{return dt.Data[i].Calendar>="2015-11-28"});
	fmt.Println(dt.Data[pk]);
}
