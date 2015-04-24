/*
接收支付平台返回的支付处理结果包，把流水号返回给支付平台，以便不再收到此包
*/
package main 

import (
 "log"
 "net/http"
 "io/ioutil"
 "net/url"
 //"strings"
 "fmt"
)

type Ret struct{
	Version, MerId, PayMoney, CardMoney, OrderId, 
	PayResult, PrivateField, PayDetails, Md5String, ErrCode,
	SignString, SzfOrderNo string;
};

func main() {
	/// 数据库连接初始化
	// var err error;
	// adm.Dbp, err = sql.Open("mysql", "bonly@tcp(127.0.0.1:3306)/moudao?charset=utf8");
	// CheckError(err);
	// adm.Dbp.SetMaxIdleConns(10);
	// defer adm.Dbp.Close();

	http.HandleFunc("/main", proc_ret);
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatal("ListenAndServe: ", err);
	}
}

func proc_ret(w http.ResponseWriter, r *http.Request){
	log.Print("req:", r.Method); //获取请求的方法
	//log.Printf("head: ", r.Header);
	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		log.Panic(err);
		return;
	}
	//log.Println(string(body));	

	//val := url.QueryEscape(string(body));
	//log.Println(val);
	val, err := url.ParseQuery(string(body));
	if err != nil{
         log.Panic(err);
         return;
	}

	ret := Ret{fmt.Sprintf("%s",val["version"]), 
	           fmt.Sprintf("%s",val["merId"]), 
	           fmt.Sprintf("%s",val["payMoney"]), 
	           "",//fmt.Sprintf("%s",val["cardMoney"]), 
	           fmt.Sprintf("%s",val["orderId"]), 
	           fmt.Sprintf("%s",val["payResult"]), 
	           fmt.Sprintf("%s",val["privateField"]), 
	           fmt.Sprintf("%s",val["payDetails"]), 
	           fmt.Sprintf("%s",val["md5String"]), 
	           fmt.Sprintf("%s",val["errcode"]),
	           fmt.Sprintf("%s",val["signString"]),
	           fmt.Sprintf("%s",val["szfOrderNo"]),
	       };

	log.Printf(`req: version=%s
		       merId=%s
		       payMoney=%s
		       cardMoney=%s
		       orderId=%s
		       payResult=%s
		       privateField=%s
		       payDetails=%s
		       md5String=%s
		       errCode=%s
		       signString=%s
		       szfOrderNo=%s`,
		       ret.Version,ret.MerId,ret.PayMoney,ret.CardMoney,ret.OrderId,ret.PayResult,ret.PrivateField,
		       ret.PayDetails,ret.Md5String,ret.ErrCode,ret.SignString,ret.SzfOrderNo);

    w.Write([]byte(ret.OrderId));
}
