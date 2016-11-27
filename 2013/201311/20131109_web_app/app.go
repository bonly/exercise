package main 

import (
"log"
"fmt"
"net/http"
"golang.org/x/net/websocket"
"open"
_ "github.com/go-sql-driver/mysql"
"database/sql"
"errors"
"regexp"
"strconv"
)

var gdb *sql.DB;
var strdb string;

func main(){
	http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir("."))));
	http.Handle("/Main", websocket.Handler(Main));

	go open.Run("http://127.0.0.1:9998/");

	err := http.ListenAndServe(":9998", nil);
	if err != nil{
		panic(err.Error());
	}
}

type Head struct{
	Cmd string;
};

type Ret struct{
	Cmd string `json:"Cmd"`;
	Ret string `json:"Code"`;
	Msg string `json:"Msg"`;
};

func Main(ws *websocket.Conn){
	log.Println(fmt.Sprintf("Get a connected from %s", "a"));

	for {
		//取消息头
		var head Head;
		err := websocket.JSON.Receive(ws, &head);
		log.Println(fmt.Sprintf("Head: %s", head));
		if err != nil{
			log.Println(err);
			return;
		}

		switch(head.Cmd){
		case "Database":
			Cmd_database(ws);
			break;
		case "Modify_Stat":
			Cmd_modify_stat(ws);
			break;
		}
	}
}

type CDB struct{
	Db string `json:"Db"`;
};

type RDB struct{
	Cmd string `json:"Cmd"`;
	Ret string `json:"Code"`;
	Msg string `json:"Msg"`;
};

func Cmd_database(ws *websocket.Conn){
	var db CDB;
	err := websocket.JSON.Receive(ws, &db);
	if err != nil{
		log.Println(err);
		return;
	}

	var rdb RDB;
	rdb.Cmd = "RDB";
	strdb = db.Db;

	gdb, err = sql.Open("mysql", strdb);
	if err != nil{
		log.Println(err);
		rdb.Ret = "-1";
		rdb.Msg = err.Error();
		websocket.JSON.Send(ws, &rdb);
		return;
	}
	gdb.Close();

	rdb.Ret = "0";
	rdb.Msg = "数据库连接测试成功";

	websocket.JSON.Send(ws, &rdb);
}

type CModify_Stat struct{
	Room_id,
	Month,
	Stat,
	Day string;
};

func Cmd_modify_stat(ws *websocket.Conn){
	var ms CModify_Stat;
	err := websocket.JSON.Receive(ws, &ms);
	if err != nil{
		log.Println(err);
		return;
	}

	var ret Ret;
	ret.Cmd = "RModify_Stat";

	ret.Ret, ret.Msg, err = Modi_data(ms.Room_id, ms.Month, ms.Stat, ms.Day); //处理更新

	websocket.JSON.Send(ws, &ret);
}

//////////////////////////
type Calc struct{
	Calendar_id,
	Room_id,
	Period,
	Stat string;
};

type Status struct{
	Flag int64;
	Price int64;
};

func Select(lst *[]Calc, room_id string, period string)(rw string, aff int, err error){
	sql := fmt.Sprintf(`
		select calendar_id,room_id,period,stat from xb_calendar
			where room_id=? and period=? `);
				
	rows, err := gdb.Query(sql, room_id, period);
	if err != nil{
		log.Println("select xb_calendar ", ": ", err);
		return "", 0, err;
	}
	defer rows.Close();

	cnt := 0;
	for rows.Next(){
		var info Calc;
		if err := rows.Scan(&info.Calendar_id, &info.Room_id, &info.Period, &info.Stat); err != nil{
			log.Println(" Row err: ", err);
			continue;
		}
		rw = fmt.Sprintf("适配: %+v", info);
		log.Println(rw);
		*lst = append(*lst, info);
		cnt++;
	}
	if err := rows.Err(); err !=nil{
		log.Println(" record err: ",err);
		err = err;
	}
	// log.Println("cnt: ", cnt);
	return rw, cnt, err;
}

func Update(tbl string, key string, key_val string, con string)(aff int, err error){
	qry := fmt.Sprintf("update %s set %s where %s = ?", tbl, con, key);
	log.Println(qry, " :key=", key_val);

	result, err := gdb.Exec(qry, key_val);
	if err != nil{
		log.Println("Update failed: ", err);
		return -1,err;
	}
	
	num, err := result.RowsAffected();
	if err != nil{
		log.Println("RowsAffected faild: ", err);
		return -1, err;
	}
	if num > 0{
		log.Println("Update rows: ", num);
		err = nil;
	}else {
		err_msg := "Record not found or had update before "+ key + "=" + key_val;
		log.Println(err_msg);
		err = errors.New(err_msg);
	}
	return int(num), err;
}

//构造字段更新
func Gen_key_value(field ...string)(con string){
	for idx, str := range field{
		if idx%2==0{
			if idx == 0{
				con = str;
			}else{
				con = con + "," + str;
			}
		}else{
			con = con + "='" + str + "'";
		}
	}
	return con;
}


func Modi_data(room_id string, month string, stat string, day string)(code string, msg string, err error){
    /// 数据库连接初始化
    // gdb, err = sql.Open("mysql", "db_admin:db_admin2015@tcp(120.25.106.243:3306)/xbed_service?charset=utf8");
    // gdb, err = sql.Open("mysql", "db_writer:XqH/a5aOnAlw@tcp(112.74.195.114:3306)/xbed_service?charset=utf8");
    code = "-1";
    gdb, err = sql.Open("mysql", strdb);
    if err != nil{
    	log.Println(err);
    	return;
    }
    gdb.SetMaxIdleConns(5);
    defer gdb.Close();	
    log.Println("连接数据库成功");

    var lst []Calc;
    rw, cnt, err := Select(&lst, room_id, month); //取记录
    if cnt != 1 || err != nil{
    	log.Println("只处理单条记录 ", err);
    	msg = "只允许处理单条记录";
    	return;
    }

    var chg []*Status;
    for _, cal := range lst{
    	//fmt.Println(i, cal);
    	reg := regexp.MustCompile(`\d\|\d*`);
		for _, mt := range (reg.FindAllStringIndex(cal.Stat, -1)){
			var val Status;
			fmt.Sscanf(cal.Stat[mt[0]:mt[1]], "%d|%d", &val.Flag, &val.Price);
			// fmt.Printf("flag: %d    price: %d\n", val.Flag, val.Price);
			chg = append(chg, &val);
		}
    }

	istat, err := strconv.ParseInt(stat, 10, 16);
	if err != nil{
		log.Println("stat 格式错误");
		return;
	}

	iday, err := strconv.ParseInt(day, 10, 16);
	if err != nil{
		log.Println("day 格式错误");
		return;
	}

    chg[iday-1].Flag = istat;
    
    var cal_ret string;
	for i, calc := range chg{
		// fmt.Printf("flag: %d    price: %d\n", calc.Flag, calc.Price);
		if i == 0{
			cal_ret = fmt.Sprintf("%d|%d", calc.Flag, calc.Price);
		}else{
			cal_ret = cal_ret + fmt.Sprintf(",%d|%d", calc.Flag, calc.Price);
		}
	}    
	// fmt.Println(cal_ret);

	cnt, err = Update("xb_calendar", "calendar_id", lst[0].Calendar_id, 
		Gen_key_value("stat", cal_ret));	

	if err == nil{
		code = "0";
		msg = fmt.Sprintf("%v 更新为:%s", rw, cal_ret);
	}else{
		msg = err.Error();
	}
	return;
}
