/*
从 Yahoo 获取
历史数据
HTML Output: http://finance.yahoo.com/q/hp?s=300072.sz&d=7&e=23&f=2010&a=5&b=11&c=2010
CSV Output: http://ichart.finance.yahoo.com/table.csv?s=300072.sz&d=7&e=23&f=2010&a=5&b=11&c=2010
由于历史原因，也可以用地址 http://table.finance.yahoo.com/table.csv
http://ichart.finance.yahoo.com/table.csv?s=AAPL&c=1962
s: 股票代码 (e.g. 002036.SZ 300072.SZ 600036.SS 等)
c-a-b: 起始日期年、月、日 (月份的起始索引为0) 2010-5-11 = 2010年6月11日
f-d-e: 结束日期年、月、日 (月份的起始索引为0) 2010-7-23 = 2010年8月23日
g: 时间周期。d=每日，w=每周，m=每月，v=只返回除权数据
省略所有参数，只制定股票代码时，返回所有历史数据

“实时”数据
http://download.finance.yahoo.com/d/quotes.csv?s=300072.SZ+600036.SS&f=spl1d1t1c1ohg (f: 指定返回数据的格式)
s: 股票代码
l: 最后成交价格和时间 。l1: 最后成交价格；d1: 最后交易日期；t1: 最后交易时间
p: 昨日收盘价；o: 开盘价；h: 最高价；g: 最低价；c1: 涨幅
v: 成交量；a2: 平均每日成交量
j: 52周最低价；j5: 52周价格变动 j6: 52周涨幅 （相对最低价）
k: 52周最高价；k4: 52周价格变动 k5: 52周涨幅 （相对最高价）
m3: MA50；m4: MA20；m5:
*/
package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"strings"
	"regexp"
	"strconv"
	"bytes"
)

type Data struct{
	Code  string;
	Dt    int64;
	Open  float64;
	High  float64;
	Low   float64;
	Close float64;
	Volume float64;
	Adj_Close float64;
};

var db *sql.DB;

var quate = []string{"A1507","AG1507"};

func main(){
   var err error;
   db, err = sql.Open("sqlite3", "./quant.db");
   if err != nil{
	   log.Fatal(err);
   }
   defer db.Close();
   
   createDb();
   
   dat := get_yahoo_year();
   if dat == nil{
   	return;
   }

   buf := bytes.NewBuffer(dat);
   for {
   	line, err := buf.ReadString('\n');
   	if err != nil{
   		fmt.Println(err);
   		break;
   	}
   	split_csv_data(line,"GOOG");
   }

   // for _, ig := range quate{
   // 		str := get_sina(ig); //网上取回数据

   // 		split_data(str, ig);
   // }
}

func split_data(str string, key string){
   reg, err := regexp.Compile(`.\d{4}-\d{2}-\d{2}.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*.`);
   if err != nil{
	   log.Fatal(err);
   }
   vt, err := regexp.Compile(`\d{4}-\d{2}-\d{2}|\d+\.\d+|\d+`);
   if err != nil{
	   log.Fatal(err);
   }
   
   // dt, err := regexp.Compile(`(?P<year>\d{4})-(?P<mon>\d{2})-(?P<day>\d{2})`);
   // if err != nil{
	//    log.Fatal(err);
   // }
   // 用dt.FindStringSubmatch 可以分段取出各个字段
   
   //分离出每天一行的数据
   for _, v := range(reg.FindAllStringSubmatch(str,-1)) {
	    var stock Data;
	    stock.Code = key;
	    // fmt.Println(idx, ":", v);
	    //分离出时间和四个浮点数及1个整数值	 
		for bdx, x := range(vt.FindAllStringSubmatch(v[0], -1)){
			// fmt.Println(bdx, ":", x);
		    switch(bdx){
				case 0:
					stock.Dt, _ = strconv.ParseInt(strings.Replace(x[0], "-", "", -1), 10, 64);
					break;
			    case 1:
			     	stock.Open, err = strconv.ParseFloat(x[0], 10);
				 	break;
				case 2:
				    stock.High, err = strconv.ParseFloat(x[0], 10);
					break;
			    case 3:
				    stock.Low, err = strconv.ParseFloat(x[0], 10);
					break;
				case 4:
					stock.Close, err = strconv.ParseFloat(x[0], 10);
					break;
				case 5:
					stock.Volume, err = strconv.ParseFloat(x[0], 10);
					break;
	        }

		}	     

	   // fmt.Println(stock);
	   //插入到数据库
	   insertData(stock);
   }
}

func insertData(dt Data){
	stmt, err := db.Prepare(`
		insert into history(Code, Dt, Open, High, Low, Close, Volume, Adj_Close)
		values (?, ?, ?, ?, ?, ?, ?, ?)`);
	if err != nil{
		log.Fatal(err);
	}
	_, err = stmt.Exec(dt.Code, dt.Dt, 
						dt.Open, dt.High,
						dt.Low, dt.Close,
						dt.Volume, dt.Adj_Close);
    if err != nil{
		log.Fatal(err);
	}
}

func createDb(){
	_, err := db.Exec(`
		create table if not exists history(
			pk INTEGER PRIMARY KEY,
			Code varchar(22),
			Dt   int(22),
			Open float(22),
			High float(22),
			Low  float(22),
			Close float(22),
			Volume float(22),
			Adj_Close float(22)
		)`);
	if err != nil{
		log.Fatal(err);
	}
}

func get_yahoo(){
    resp, err := http.Get("http://ichart.finance.yahoo.com/table.csv?s=300072.sz");
    if err != nil {
        log.Println("http get: ",err);
    }
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
    }
 
    fmt.Println(string(body));	
}

func get_sina(key string) string{
    resp, err := http.Get("http://stock2.finance.sina.com.cn/futures/api/json.php/IndexService.getInnerFuturesDailyKLine?symbol="+key);
    if err != nil {
        log.Println("http get: ",err);
    }
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
    }
 
    fmt.Println(string(body));	
	return string(body);
}

func get_yahoo_year()([]byte){
    resp, err := http.Get("http://ichart.finance.yahoo.com/table.csv?s=GOOG&c=2015");
    if err != nil {
        log.Println("http get: ",err);
        return nil;
    }
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
        return nil;
    }
 
    // fmt.Println(string(body));	
    return body;
}

func split_csv_data(str string, key string){
   reg, err := regexp.Compile(`\d{4}-\d{2}-\d{2}.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*.\d{2}\.\d*`);
   if err != nil{
	   log.Fatal(err);
   }
   vt, err := regexp.Compile(`\d{4}-\d{2}-\d{2}|\d+\.\d+|\d+`);
   if err != nil{
	   log.Fatal(err);
   }
   
   // dt, err := regexp.Compile(`(?P<year>\d{4})-(?P<mon>\d{2})-(?P<day>\d{2})`);
   // if err != nil{
	//    log.Fatal(err);
   // }
   // 用dt.FindStringSubmatch 可以分段取出各个字段
   
   //分离出每天一行的数据
   for _, v := range(reg.FindAllStringSubmatch(str,-1)) {
	    var stock Data;
	    stock.Code = key;
	    // fmt.Println(idx, ":", v);
	    //分离出时间和四个浮点数及1个整数值	 
		for bdx, x := range(vt.FindAllStringSubmatch(v[0], -1)){
			// fmt.Println(bdx, ":", x);
		    switch(bdx){
				case 0:
					stock.Dt, _ = strconv.ParseInt(strings.Replace(x[0], "-", "", -1), 10, 64);
					break;
			    case 1:
			     	stock.Open, err = strconv.ParseFloat(x[0], 10);
				 	break;
				case 2:
				    stock.High, err = strconv.ParseFloat(x[0], 10);
					break;
			    case 3:
				    stock.Low, err = strconv.ParseFloat(x[0], 10);
					break;
				case 4:
					stock.Close, err = strconv.ParseFloat(x[0], 10);
					break;
				case 5:
					stock.Volume, err = strconv.ParseFloat(x[0], 10);
					break;
				case 6:
					stock.Adj_Close, err = strconv.ParseFloat(x[0], 10);
					break;
	        }

		}	     

	   // fmt.Println(stock);
	   //插入到数据库
	   insertData(stock);
   }
}

//go install -a github.com/mattn/go-sqlite3 会把.a编译好放到库中，不需每次重编译
//go build -i 将更快