package main

import (
"fmt"
"encoding/json"
)

/*
template sms
*/
type Tpl_Msg struct{
	ToUser string `json:"touser"`;
	Template_id string `json:"template_id"`;
	Url string `json:"url"`;
	TopColor string `json:"topcolor"`;
	Data interface{} `json:"data"`;
};

type Order_success struct{
	OrderID string `json:"orderID"`;
	OrderMoneySum string `json:"orderMoneySum"`;
};

func main(){
	var msg Tpl_Msg;
	data := Order_success{
		"12345678",
		"30.00元",
	};	
	msg.Data = data;
	bt, err := json.MarshalIndent(msg, " ", " ");
	if err != nil{
		fmt.Println(err);
		return;
	}

	fmt.Println(string(bt));
}