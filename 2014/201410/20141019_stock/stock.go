package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"time"
	"regexp"
	"strconv"
	"bytes"	
)

import (
"github.com/gonum/plot"
"github.com/gonum/plot/plotter"
// "github.com/gonum/plot/plotutil" //method 1
"github.com/gonum/plot/vg"

//method 2
"github.com/gonum/plot/vg/draw"
"image/color"
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

   tr, err := plot.New();
   if err != nil{
   		panic(err);
   }

   tr.Title.Text =  "GOOG";
   tr.X.Label.Text = "日期";
   tr.Y.Label.Text = "Price";
   tr.X.Tick.Marker = plot.UnixTimeTicks{Format: "2006-01-02"};

   var ln plotter.XYs;

   buf := bytes.NewBuffer(dat);
   for {
	   	line, err := buf.ReadString('\n');
	   	if err != nil{
	   		fmt.Println(err);
	   		break;
	   	}
	   	split_csv_data(line,"GOOG", &ln);
   }

   //method 2
	line, points, err := plotter.NewLinePoints(ln);
	if err != nil {
	    log.Panic(err);
	}
	line.Color = color.RGBA{G: 255, A: 255};
	points.Shape = draw.CircleGlyph{};
	points.Color = color.RGBA{R: 255, A: 255};
	tr.Add(line, points);

	//method 1
   // err = plotutil.AddLinePoints(tr, "Open", ln);
   // if err != nil{
   // 	  panic(err);
   // }

   if err := tr.Save(12*vg.Inch, 6*vg.Inch, "goog.svg"); err != nil{
   	  panic(err);
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

func split_csv_data(str string, key string, xy *plotter.XYs){
   reg, err := regexp.Compile(`\d{4}-\d{2}-\d{2}.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*.\d{2}\.\d*`);
   if err != nil{
	   log.Fatal(err);
   }
   vt, err := regexp.Compile(`\d{4}-\d{2}-\d{2}|\d+\.\d+|\d+`);
   if err != nil{
	   log.Fatal(err);
   }
   // dt, err := regexp.Compile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`);
   // if err != nil{
	  //  log.Fatal(err);
   // }
   // 用dt.FindStringSubmatch 可以分段取出各个字段   
   
   //分离出每天一行的数据
   for _, v := range(reg.FindAllStringSubmatch(str,-1)) {
	    var stock Data;
	    stock.Code = key;
	    // fmt.Println(idx, ":", v);
	    //分离出时间和四个浮点数及1个整数值	 
	    var date int64;
		for bdx, x := range(vt.FindAllStringSubmatch(v[0], -1)){
			// fmt.Println(bdx, ":", x);
		    switch(bdx){
				case 0:
					// fmt.Println(x[0]);
					var year, month, day, hour, min, sec, nsec int;
					_, err := fmt.Sscanf(x[0], "%d-%d-%d", &year, &month, &day);
					if err != nil{
						fmt.Println(err);
						continue;
					}
					date = time.Date(year, (time.Month)(month), day, hour, min, sec, nsec, time.UTC).Unix();
					stock.Dt = date;
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
	   *xy = append(*xy, plotter.XYs{{X:float64(date), Y:stock.Open}}[0]);
   }
}
