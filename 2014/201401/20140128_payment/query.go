package main 


import (
"fmt"
// "io/ioutil"
// "net/http"
"golang.org/x/net/websocket"
// "open"
"glog"
_ "github.com/go-sql-driver/mysql"
"database/sql"
"encoding/json"
)

type Q_Query struct{
	Oms_order_id string;
};

//order_no,lodger_name,lodger_mobile,platform_sn,p.pay_type,pay_money/100.0,pay_time,addition_data
type R_Query struct{
	Ret;
	Data []Payment;
};

type Payment struct{
	Order_no string `db:"Order_no"`;
	Lodger_name string `db:"Lodger_name"`;
	Lodger_mobile string `db:"Lodger_mobile"`;
	Platform_sn string `db:"Platform_sn"`;
	Pay_type string `db:"Pay_type"`;
	Pay_money string `db:"Pay_money"`;
	Pay_time string `db:"Pay_time"`;
	Addition_data string `db:"Addition_data"`;
};

func Cmd_Query(ws *websocket.Conn, it interface{}){
	ret := it.(*R_Query);
	ret.Ret.Cmd = "RQuery";
	var qry Q_Query;
	err := websocket.JSON.Receive(ws, &qry);
	if err != nil{
		glog.Info("qry body error");
		ret.Ret.Ret = "1";
		ret.Ret.Msg = "body error";
		return;
	}

	glog.Info(qry);
	ret.Ret.Ret = "0";
	ret.Ret.Msg = "OK";

	//查询
	Get_Payment(qry.Oms_order_id, ret);

	js, err := json.MarshalIndent(ret, " ", " ");
	if err != nil{
		glog.Info("encode json: ", err);
		return;
	}

	glog.Info("send: ",string(js));	
}

func Get_Payment(order_id string, ret *R_Query){
	// glog.Info(fmt.Sprintf("query: %s", order_id));
	qry :=`select Order_no,Lodger_name,Lodger_mobile,Platform_sn,p.pay_type Pay_type,
		                           pay_money/100.0 Pay_money, Pay_time, Addition_data
							from xb_order o, xb_payment p
							where o.order_no=p.order_id and o.order_no= ? and recv_code='SUCCESS' `;
	glog.Info(fmt.Sprintf("qry: \n %s; order_no[%s]\n", qry, order_id));
	rows, err := db.Queryx(qry, order_id);
	// defer rows.Close();
	switch{
		case err == sql.ErrNoRows:{
			glog.Info("query: ", err);
			return;
		}
		case err != nil: {
			glog.Info("query: ", err);
			return;
		}
	}
	
	for rows.Next(){
		var pay Payment;
		err := rows.StructScan(&pay);
		if err != nil{
			glog.Info("row: ", err);
			continue;
		}
		ret.Data = append(ret.Data, pay);
		glog.Info(fmt.Sprintf("%+v", pay));
	}

}