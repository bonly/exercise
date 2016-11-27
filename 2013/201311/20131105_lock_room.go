package main 

import (
"log"
"fmt"
"database/sql"
_ "github.com/go-sql-driver/mysql"
"os"
"regexp"
"errors"
"strconv"
)

var Db_r *sql.DB;

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

func main(){
	if len(os.Args) < 4{
		log.Println(os.Args[0], " room_id ", " period ", " day ", " stat ");
		return;
	}
	day, err := strconv.ParseInt(os.Args[3], 10, 16);
	if err != nil{
		log.Println("day 格式错误");
		return;
	}
	day = day - 1;
	stat, err := strconv.ParseInt(os.Args[4], 10, 16);
	if err != nil{
		log.Println("stat 格式错误");
		return;
	}	

    /// 数据库连接初始化
    // Db_r, err = sql.Open("mysql", "db_admin:db_admin2015@tcp(120.25.106.243:3306)/xbed_service?charset=utf8");
    Db_r, err = sql.Open("mysql", "db_writer:XqH/a5aOnAlw@tcp(112.74.195.114:3306)/xbed_service?charset=utf8");
    if err != nil{
    	log.Println(err);
    	return;
    }
    Db_r.SetMaxIdleConns(5);
    defer Db_r.Close();	

    var lst []Calc;
    cnt, err := Select(&lst, os.Args[1], os.Args[2]); //取记录
    if cnt != 1 || err != nil{
    	log.Println("只处理单条记录 ", err);
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

    chg[day].Flag = stat;
    
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
}

func Select(lst *[]Calc, room_id string, period string)(aff int, err error){
	sql := fmt.Sprintf(`
		select calendar_id,room_id,period,stat from xb_calendar
			where room_id=? and period=? `);
				
	rows, err := Db_r.Query(sql, room_id, period);
	if err != nil{
		log.Println("select xb_calendar ", ": ", err);
		return 0, err;
	}
	defer rows.Close();

	cnt := 0;
	for rows.Next(){
		var info Calc;
		if err := rows.Scan(&info.Calendar_id, &info.Room_id, &info.Period, &info.Stat); err != nil{
			log.Println(" Row err: ", err);
			continue;
		}
		log.Println(fmt.Sprintf("适配: %+v", info));
		*lst = append(*lst, info);
		cnt++;
	}
	if err := rows.Err(); err !=nil{
		log.Println(" record err: ",err);
		err = err;
	}
	// log.Println("cnt: ", cnt);
	return cnt, err;
}

func Update(tbl string, key string, key_val string, con string)(aff int, err error){
	qry := fmt.Sprintf("update %s set %s where %s = ?", tbl, con, key);
	log.Println(qry, " :key=", key_val);

	result, err := Db_r.Exec(qry, key_val);
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
