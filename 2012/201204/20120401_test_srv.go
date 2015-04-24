package main

import (
"log"
"net/http"
"io/ioutil"
"encoding/json"
)

type R_ret struct{
Ret string `json:"ret"`;
Msg string `json:"msg"`;
User_id string `json:"user_id"`;
};

func main(){
	http.HandleFunc("/", Ret);
	err := http.ListenAndServe(":13999", nil)
	if err != nil{
		log.Println(err.Error());
	}
}

func Ret(rw http.ResponseWriter, qry *http.Request){
	log.Println("get a request");
	defer log.Println("process end");

    log.Println(qry.Method);
    
	body, err := ioutil.ReadAll(qry.Body);
	if err != nil{
		log.Println(err.Error());
	}
	log.Println(string(body));

    var rep R_ret;
    rep.Ret = "0";
    rep.Msg = "OK";
    rep.User_id = "1001";
	back, err := json.Marshal(rep);

	rw.Write([]byte(back));
}

