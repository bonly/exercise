/*
auth: bonly
create: 2015.12.18
*/

package main 

import (
"net/http"
"log"
"github.com/jmoiron/sqlx"
_ "github.com/go-sql-driver/mysql"
"encoding/json"
"io/ioutil"
"fmt"
"database/sql"
)

/*
CREATE TABLE `analysis_room_days` (
  `id` varchar(255) NOT NULL COMMENT '主键',
  `room_id` int(50) DEFAULT NULL COMMENT '房间id',
  `th_date` char(10) DEFAULT NULL COMMENT '日结时间',
  `future_sale` int(1) DEFAULT '0' COMMENT '是否可以远期销售？ 0 不可远期销售  1 可远期销售',
  `online` int(1) DEFAULT '0' COMMENT '是否上线？  0未上线  1已上线',
  `stop` int(1) DEFAULT '0' COMMENT '是否停用？ 0未停用  1已停用',
  `over_night` int(1) DEFAULT '0' COMMENT '是否过夜？  0未过夜  1已过夜',
  `sold` int(1) DEFAULT '0' COMMENT '是否出租？ 0代表未出租房间   其它数字代表出租次数',
  `batch` varchar(255) DEFAULT NULL COMMENT '批次',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

var db *sqlx.DB;

// type Room_lst struct{
// 	Lst []Room_days;
// };

type Room_days struct{
	Room_id string `db:"Room_id"`;
	Future_sale string `db:"Future_sale"`;
	Online string `db:"Online"`;
	Stop string `db:"Stop"`;
	Over_night string `db:"Over_night"`;
	Sold string `db:"Sold"`;
};

type Q_Room_days struct{
	Date string;
	Room_id string;
};

type Order_days struct{
	Room_id string  `db:"Room_id"`;
	Real_price string `db:"Real_price"`;
	Discount string `db:"Discount"`;
	Income string `db:"Income"`;
	Channel_name string `db:"Channel_name"`;
	Commission string `db:"Commission"`;
};

func Get_order_of_room(rw http.ResponseWriter, qry *http.Request){
	log.Println("get a order of room query");
	// rw.Header.Set("Content-Type","application/x-www-form-urlencoded");
	// rw.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");

	body, err := ioutil.ReadAll(qry.Body);
	if err != nil{
		log.Println(err);
		return;
	}
	log.Println(fmt.Sprintf("recv: %s", string(body)));

	var dt Q_Room_days;
	err = json.Unmarshal(body, &dt);
	if err != nil{
		log.Println(err);
		return;
	}

	var lst []Order_days;
	rows, err := db.Queryx(`select Room_id, Real_price, Discount, Income, Channel_name, Commission
	                 from analysis_order_days
					 where room_id = ? and th_date = ? `, dt.Room_id, dt.Date);
	defer rows.Close();
	switch{
		case err == sql.ErrNoRows:{
			log.Println("query: ", err);
			return;
		}
		case err != nil: {
			log.Println("query: ", err);
			return;
		}
	}

	for rows.Next(){
		var room Order_days;
		err := rows.StructScan(&room);
		if err != nil{
			log.Println("row: ", err);
			continue;
		}
		lst = append(lst, room);
	}

	js, err := json.MarshalIndent(lst, " ", " ");
	if err != nil{
		log.Println("encode json: ", err);
		return;
	}

	log.Println("send: ",string(js));
	rw.Write(js);
}

func Get_day_of_room(rw http.ResponseWriter, qry *http.Request){
	log.Println("get a day of room query");
	// rw.Header.Set("Content-Type","application/x-www-form-urlencoded");
	// rw.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");

	body, err := ioutil.ReadAll(qry.Body);
	if err != nil{
		log.Println(err);
		return;
	}
	log.Println(fmt.Sprintf("recv: %s", string(body)));

	var dt Q_Room_days;
	err = json.Unmarshal(body, &dt);
	if err != nil{
		log.Println(err);
		return;
	}

	var lst []Room_days;
	rows, err := db.Queryx(`select Room_id, Future_sale, Online, Stop Over_night, Sold from analysis_room_days
					 where room_id = ? and th_date = ? `, dt.Room_id, dt.Date);
	defer rows.Close();
	switch{
		case err == sql.ErrNoRows:{
			log.Println("query: ", err);
			return;
		}
		case err != nil: {
			log.Println("query: ", err);
			return;
		}
	}

	for rows.Next(){
		var room Room_days;
		err := rows.StructScan(&room);
		if err != nil{
			log.Println("row: ", err);
			continue;
		}
		lst = append(lst, room);
	}

	js, err := json.MarshalIndent(lst, " ", " ");
	if err != nil{
		log.Println("encode json: ", err);
		return;
	}

	log.Println("send: ",string(js));
	rw.Write(js);
}

func main(){
	var err error;
	db, err = sqlx.Open("mysql", "db_admin:db_admin2015@tcp(120.25.106.243:3306)/xbed_analysis?charset=utf8");
	if err != nil{
		log.Println(err);
		return;
	}

	http.HandleFunc("/room_data", Get_day_of_room);
	http.HandleFunc("/order_data", Get_order_of_room);
	err = http.ListenAndServe("0.0.0.0:7777", nil);
	defer func(){
		if err := recover(); err != nil{
			log.Println(err);
			return;
		}
	}();
}