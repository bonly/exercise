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

import (
"github.com/vdobler/chart"
// "github.com/vdobler/chart/txtg"
"github.com/ajstarks/svgo"
"github.com/vdobler/chart/svgg"
"image/color"
// "math/rand"
"os"
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

   tr := chart.ScatterChart{Title: "GOOG"};
   var x []float64;
   var y []float64;   

   buf := bytes.NewBuffer(dat);
   idx := 0.0;
   for {
	   	line, err := buf.ReadString('\n');
	   	if err != nil{
	   		fmt.Println(err);
	   		break;
	   	}
	   	split_csv_data(line,"GOOG", &idx, &x, &y);
   }
   tr.AddDataPair("Open", x, y, chart.PlotStyleLinesPoints, chart.AutoStyle(1, true));
   
   svg_file, _ := os.Create("goog.svg");
   defer svg_file.Close();
   svg_grap := svg.New(svg_file);
   svg_grap.Start(500, 500);
   defer svg_grap.End();

   svg_grap.Title("Price");
   svg_grap.Rect(0, 0, 500, 1500, "fill: #ffffff");
   svgr := svgg.AddTo(svg_grap, 10, 10, 500, 500, "", 12, color.RGBA{0xff, 0xff, 0xff, 0xff});
   tr.Plot(svgr);   
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

func split_csv_data(str string, key string, idx *float64, x *[]float64, y *[]float64){
   reg, err := regexp.Compile(`\d{4}-\d{2}-\d{2}.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*\.\d*.{3}\d*.\d{2}\.\d*`);
   if err != nil{
	   log.Fatal(err);
   }
   vt, err := regexp.Compile(`\d{4}-\d{2}-\d{2}|\d+\.\d+|\d+`);
   if err != nil{
	   log.Fatal(err);
   }
   
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
   	   *x = append(*x, *idx);
	   *y = append(*y, stock.Open);
	   *idx = *idx + 1;
   }
}
