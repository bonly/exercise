/*
把mysql中的json转换成记录
*/
package main

import (
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/go-sql-driver/mysql"
	"flag"
	"fmt"
	"encoding/json"
	"strings"
)

var db *sqlx.DB;

var db_connect = flag.String("d", 
	// "root:techappen@tcp(192.168.1.104:3306)/techappen?charset=utf8",
	"root:techappen@tcp(139.199.203.251:3306)/techappen?charset=utf8",
	"database");

func init(){
	flag.Parse();
}

func main(){
	var err error;
	db, err = sqlx.Open("mysql", *db_connect);
	if err != nil{
		log.Fatalf("db: %v\n", err);
	}

	//取旧用户资料
	users := []string{};
	if err = Get_User(&users); err!=nil{
		log.Fatalf("user: %v\n", err);
	}

	for idx, user := range users{
		log.Printf("process %d user[%s]\n", idx, user);
		pairs := []*Pair{};
		//取旧表数据
		if err = AllPair("scene", user+"SCENE", &pairs); err != nil{
			log.Printf("scene: %v\n", err);
			return;
		}

		if len(pairs) <= 0{
			continue;
		}

		//检查表是否存在
		//不存在则创建新表
		if exists, _ := Chk("Build_"+user); !exists{
			Create("Build_"+user);	
		}
		if exists, _ := Chk("Disturb_"+user); !exists{
			Create("Disturb_"+user);	
		}

		//转换分析vdata
		var st Scene;
		if err = json.Unmarshal([]byte(pairs[0].Value), &st); err != nil{
			log.Printf("json %s: %v\n", pairs[0].Key, err);
			continue;
		}

		//遍历数据
		for idx_obj, data := range st.Value {
			js,_ := json.Marshal(data);
			fmt.Printf("v%d: %v\n", idx_obj, data);
			//分析并插入到新表中
			if strings.Contains(data.BuildId, "Tree"){
				log.Printf("树: %s\n", js);
				Set("Disturb_"+user, data.Id, string(js));
			}else if strings.Contains(data.BuildId, "Rock"){
				log.Printf("石: %s\n", js);
				Set("Disturb_"+user, data.Id, string(js));
			}else{
				log.Printf("障碍: %s\n", js);
				Set("Build_"+user, data.Id, string(js));
			}
		}
	}
}

type Pair struct{
	Key string;
	Value string;
};

type Data struct{
	Id string;
	Name string;
	BuildId string;
	Pos string;
	Belong int;
	Rotate int;
	NotEdit bool;
};

type Scene struct{
	Key []string;
	Value []Data;
};

//取旧用户资料
func Get_User(users *[]string) (err error){
	qry := fmt.Sprintf(`select user_id from user where user_name like ?`);

	res, err := db.Query(qry, "lz%");
	if err != nil{
    	log.Printf("get user all failed: %v\n", err);
		return err;
	}
	defer res.Close();
	for res.Next() {
		var value string;
        err := res.Scan(&value);
        if err != nil {
            log.Printf("%v\n", err);
			return err;
        } 
		*users = append(*users, value);
    }

	return nil;	
}

//取旧表数据
func AllPair(table string, key string, pairs *[]*Pair) (err error){
	qry := fmt.Sprintf("select vkey,vdata from %s where vkey=?", table);

	res, err := db.Query(qry, key);
	if err != nil{
		if driverErr, ok := err.(*mysql.MySQLError); ok{
			if driverErr.Number == 1146 { //表不存在
				log.Printf("[%s]表不存在\n", table);
				return fmt.Errorf("表不存在\n");
			
		    }
		}else{
	    	log.Printf("get table all failed: %v\n", err);
	    }
		return err;
	}
	defer res.Close();
	for res.Next() {
		value := Pair{};
        err := res.Scan(&value.Key, &value.Value);
        if err != nil {
            log.Printf("%v\n", err);
			return err;
        } 
		*pairs = append(*pairs, &value);
    }

	return nil;	
}

/*
创建
*/
func Create(table string)(err error){
	log.Printf("创建表[%s]", table);
	qry := fmt.Sprintf(`
CREATE TABLE %s (
  vkey varchar(100) NOT NULL DEFAULT 'none' COMMENT 'key',
  vdata mediumtext COMMENT '用户数据',
  PRIMARY KEY (vkey)
)
`, table);

	_, err = db.Exec(qry);
	if err != nil{
		log.Printf("create table failed: %v\n", err);
		return err;
	}
	return nil;
}

/*
写入
*/
func Set(table string, key string, value string)(err error){
	qry := fmt.Sprintf(`INSERT INTO %s (vkey,vdata)
                        VALUES (?, ?) ON DUPLICATE KEY UPDATE vdata=?`, table);

	res, err := db.Exec(qry, key, value, value);
	if err != nil{
		if driverErr, ok := err.(*mysql.MySQLError); ok{
			if driverErr.Number == 1146{ //表不存在
				log.Printf("[%s]表不存在\n",table);
				return fmt.Errorf("表不存在\n");
			}
		}
		log.Printf("ins key failed: %v\n", err);
		return err;
	}
	affr, _ := res.RowsAffected();
	if affr != 1{
		log.Printf("Affect row: %v\n", affr);
	}
	return nil;
}

//检查表是否存在
func Chk(table string) (exists bool, err error){
	qry := fmt.Sprintf("Select 1 from %s LIMIT 1", table);

	res, err := db.Query(qry);
	if err != nil{
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			// if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
			if driverErr.Number == 1146 { //表不存在
				log.Printf("[%s]表不存在\n", table);
				return false, nil;
			}
		}else{
		     log.Printf("get key failed: %v\n", err);
		}
		return true, err;
	}
	defer res.Close();

	return true, nil;
}