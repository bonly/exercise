package main 

import (
"log"
"net/http"
"os"
"io"
"encoding/json"
"fmt"
// "strings"
)

type Card_info struct{
	Card_id string;
	Name string;
	Sex string;
	Nation string;
	Birthday string;
	Address string;
	Mac_id string;
	Ver string;
};

type Ret struct{
	Ret string;
	Ret_msg string;
};

func main(){
	http.HandleFunc("/idcard", process_idcard);

	err := http.ListenAndServe(":6666", nil);
	if err != nil{
		log.Fatal(err);
	}
}

func process_idcard(rw http.ResponseWriter, qry *http.Request){
	log.Println("=========== get a id card info =============");
	defer func(){
		log.Println("============ end process id card info ===========");
	}();

	ret, err := bus(rw, qry);

	str, err := json.Marshal(ret);
	if err != nil{
		log.Println("json: ", err);
		return;
	}

	rw.Write(str);
	log.Println("return: ", string(str));
}

func bus(rw http.ResponseWriter, qry *http.Request) (ret Ret, err error){
	ret.Ret = "0";
	ret.Ret_msg = "Success";

	if code, terr := save_pic(rw, qry); terr != nil{
		ret.Ret = code;
		ret.Ret_msg = terr.Error();
		return;
	}
	
	id_info := qry.FormValue("Card_info");
	log.Println("身份证:", id_info);

	if code, terr := bus_ver(rw, qry); terr != nil{
		ret.Ret = code;
		ret.Ret_msg = terr.Error();
		return;
	}

	verify := qry.FormValue("Verify");
	log.Println("授权码: ", verify);

	order := qry.FormValue("Order");
	log.Println("订单: ", order);

	if code, terr := bus_idcard(rw, qry); terr != nil{
		ret.Ret = code;
		ret.Ret_msg = terr.Error();
		return;
	}
	return;
}

func bus_ver(rw http.ResponseWriter, qry *http.Request) (ret string, err error){
	ver := qry.FormValue("Ver");
	log.Println("版本:", ver);	
	if ver != "V1.0" {
		ret = "200";
		err = fmt.Errorf("版本号不支持%s", ver);
		return;
	}
	return "0", nil;
}

func bus_idcard(rw http.ResponseWriter, qry *http.Request) (ret string, err error){
	return "0", nil;
}

func save_pic(rw http.ResponseWriter, qry *http.Request)(ret string, err error){
	qry.ParseMultipartForm(0);
	defer qry.MultipartForm.RemoveAll();

	fi, info, err := qry.FormFile("Pic");
	if err != nil{
		log.Println("get file: ", err);
		ret = "100";
		return;
	}
	defer fi.Close();
	log.Printf("Recieved %v", info.Filename);

	ofi, err := os.OpenFile(info.Filename, os.O_CREATE|os.O_RDWR, 0666);
	if err != nil{
		log.Println("create file: ", err);
		ret = "101";
		return;
	}
	defer ofi.Close();	

	_, err = io.Copy(ofi, fi);
	if err != nil{
		log.Println("save file: ", err);
		ret = "102";
		return;
	}

	return "0", nil;
}