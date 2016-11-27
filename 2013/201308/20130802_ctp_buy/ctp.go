package main

/*
#cgo CPPFLAGS: -g
#cgo LDFLAGS: -lthostmduserapi -lthosttraderapi -L.
#include "ctp_tg.h"
#include "ThostFtdcUserApiStruct.h"
*/
import "C"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"
	"encoding/binary"
	// "path/filepath"
	// "strings"
	"time"
	"net/http"
	"database/sql"
	// "github.com/go-sql-driver/mysql"
	_"github.com/mattn/go-sqlite3"
	"github.com/qiniu/iconv"
)

type Head struct{
	Version		[4]byte; //uint32;
	Copyright 	[64]byte;
	Symbol 		[12]byte;
	Period 		[4]byte;//int32;
	Digits 		[4]byte;//int32;
	Timesign 	[4]byte;//uint32;
	Last_sync 	[4]byte;//uint32;
	Unused		[52]byte;
};

type Bar struct{
	Ctm 		int64;//[8]byte;
	Open 		float64;//[8]byte;
	High		float64;//[8]byte;
	Low			float64;//[8]byte;
	Close		float64;//[8]byte;
	Volume		int64;//[8]byte;
	Spread		int32;//[4]byte;
	Real_volume int64;//[8]byte;
};

var Dbp *sql.DB

func main(){
	sigs := make(chan os.Signal, 1);
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);
	
	fmt.Println("begin");
	
	var err error;
	// Dbp, err = sql.Open("mysql", "bonly@tcp(127.0.0.1:3306)/quant?charset=utf8");
	Dbp, err = sql.Open("sqlite3", "./quant.db");
	if err != nil{
		log.Fatal(err);
	}
	defer Dbp.Close();
	
	_, err = Dbp.Exec(`
		create table  if not exists ThostFtdcDepthMarketData (
					TradingDay  varchar(22),
					InstrumentID varchar(22),
					ExchangeID  varchar(22),
					ExchangeInstID varchar(22),
					LastPrice  float(22),
					PreSettlementPrice float(22),
					PreClosePrice  float(22),
					PreOpenInterest  int(22),
					OpenPrice  float(22),
					HighestPrice  float(22),
					LowestPrice  float(22),
					Volume  int(22),
					Turnover  float(22),
					OpenInterest int(22),
					ClosePrice float(22),
					SettlementPrice float(22),
					UpperLimitPrice float(22),
					LowerLimitPrice float(22),
					PreDelta  float(22),
					CurrDelta  float(22),
					UpdateTime varchar(22),
					UpdateMillisec int(22),
					BidPrice1  float(22),
					BidVolume1  int(22),
					AskPrice1  float(22),
					AskVolume1  int(22),
					BidPrice2  float(22),
					BidVolume2  int(22),
					AskPrice2  float(22),
					AskVolume2	int(22),
					BidPrice3  float(22),
					BidVolume3  int(22),
					AskPrice3  float(22),
					AskVolume3  int(22),
					BidPrice4  float(22),
					BidVolume4  int(22),
					AskPrice4  float(22),
					AskVolume4  int(22),
					BidPrice5  float(22),
					BidVolume5  int(22),
					AskPrice5  float(22),
					AskVolume5  int(22),
					AveragePrice float(22),
					ActionDay  varchar(22)
		)
	`);
	if err != nil{
		log.Fatal(err);
	}
	
	go html();
	// go C.market();
	go C.trader();
	
	<-sigs;

	fmt.Println("end");
}

 
func html(){
	http.Handle("/", http.FileServer(http.Dir(".")));
	err := http.ListenAndServe(":8888", nil);
	if err != nil{
		log.Fatal(err);
	}
}

//export data2db
func data2db(da *C.struct_CThostFtdcDepthMarketDataField){
	stmt, err := Dbp.Prepare(`
		insert into ThostFtdcDepthMarketData(
					TradingDay,
					InstrumentID,
					ExchangeID,
					ExchangeInstID,
					LastPrice,
					PreSettlementPrice,
					PreClosePrice,
					PreOpenInterest,
					OpenPrice,
					HighestPrice ,
					LowestPrice ,
					Volume,
					Turnover ,
					OpenInterest ,
					ClosePrice ,
					SettlementPrice ,
					UpperLimitPrice ,
					LowerLimitPrice,
					PreDelta ,
					CurrDelta ,
					UpdateTime,
					UpdateMillisec,
					BidPrice1,
					BidVolume1 ,
					AskPrice1,
					AskVolume1 ,
					BidPrice2 ,
					BidVolume2 ,
					AskPrice2,
					AskVolume2	,
					BidPrice3,
					BidVolume3,
					AskPrice3 ,
					AskVolume3 ,
					BidPrice4 ,
					BidVolume4 ,
					AskPrice4 ,
					AskVolume4 ,
					BidPrice5 ,
					BidVolume5 ,
					AskPrice5 ,
					AskVolume5 ,
					AveragePrice ,
					ActionDay
		)values(
		  ?,?,?,?,?,?,?,?,?,?, 
		  ?,?,?,?,?,?,?,?,?,?, 
		  ?,?,?,?,?,?,?,?,?,?, 
		  ?,?,?,?,?,?,?,?,?,?, 
		  ?,?,?,?
		)
	`);

	res, err := stmt.Exec(
		C.GoString(&(da.TradingDay[0])),
		C.GoString(&(da.InstrumentID[0])),
		C.GoString(&(da.ExchangeID[0])),
		C.GoString(&(da.ExchangeInstID[0])),
		da.LastPrice,
		da.PreSettlementPrice,
		da.PreClosePrice,
		da.PreOpenInterest,
		da.OpenPrice,
		da.HighestPrice ,
		da.LowestPrice ,
		da.Volume,
		da.Turnover ,
		da.OpenInterest ,
		da.ClosePrice ,
		da.SettlementPrice ,
		da.UpperLimitPrice ,
		da.LowerLimitPrice,
		da.PreDelta ,
		da.CurrDelta ,
		C.GoString(&(da.UpdateTime[0])),
		da.UpdateMillisec,
		da.BidPrice1,
		da.BidVolume1 ,
		da.AskPrice1,
		da.AskVolume1 ,
		da.BidPrice2 ,
		da.BidVolume2 ,
		da.AskPrice2,
		da.AskVolume2	,
		da.BidPrice3,
		da.BidVolume3,
		da.AskPrice3 ,
		da.AskVolume3 ,
		da.BidPrice4 ,
		da.BidVolume4 ,
		da.AskPrice4 ,
		da.AskVolume4 ,
		da.BidPrice5 ,
		da.BidVolume5 ,
		da.AskPrice5 ,
		da.AskVolume5 ,
		da.AveragePrice ,
		C.GoString(&(da.ActionDay[0])));
		
    if err != nil{
		log.Println(err);
	}
	id, err := res.LastInsertId();
	if err != nil{
		log.Println(err);
	}
	log.Println("lastid: ", id);
}


//export data2hst
func data2hst(da *C.struct_CThostFtdcDepthMarketDataField){
	ofile := "abc.log";
	var of *os.File;
	if _, err := os.Stat(ofile); os.IsNotExist(err) {
	    fmt.Printf("create file: %s", ofile);
		of, err = os.Create(ofile);
		if err != nil{
			log.Panic(err);
		}
		of.Seek(0, os.SEEK_SET);//到头
	}else{
		of, err = os.OpenFile(ofile,os.O_WRONLY|os.O_APPEND, 777);
		if err != nil{
			log.Fatal("open file failed: ", err);
		}
		of.Seek(0, os.SEEK_END);//到尾
	}
	defer func(){
		of.Sync();
		of.Close();
	}();
	
	ret, err := of.Seek(0, os.SEEK_CUR);
	if int64(os.SEEK_SET) == ret {//在文件头
		  var head Head;
		  binary.LittleEndian.PutUint32(head.Version[:], 401);
		  copy(head.Copyright[:], "(C)opyright 2003, MetaQuotes Software Corp.");
		  // copy(head.Symbol[:], strings.TrimSuffix(filepath.Base(ifile),filepath.Ext(ifile)));
		  copy(head.Symbol[:], C.GoString(&(da.InstrumentID[0])));
		  binary.LittleEndian.PutUint32(head.Period[:], 1); //写什么值？
		  binary.LittleEndian.PutUint32(head.Digits[:], 5); //写什么值?
		  binary.LittleEndian.PutUint32(head.Timesign[:], (uint32)(time.Now().Unix()));
		  
		  if err = binary.Write(of, binary.LittleEndian, &head); err!=nil{
		    log.Panic("write head err: ", err);
		  }
	}
	
	var bar Bar;
	str := C.GoString(&(da.TradingDay[0])) + " " + 
	      C.GoString(&(da.UpdateTime[0]));
    tm, err := time.Parse("20060102 15:04:05", str);
	if err != nil{
		log.Println("time err: ", err);
	}
	bar.Ctm = tm.Unix();
	bar.Open = float64(da.OpenPrice);
	bar.High = float64(da.HighestPrice);
	bar.Low = float64(da.LowestPrice);
    bar.Close = float64(da.LastPrice);
	bar.Volume = int64(da.Volume);
	// bar.Turnover = da.Turnover;
	
	if err = binary.Write(of, binary.LittleEndian, &bar); err!=nil{
		log.Panic("write bar err: ", err);
		return;
	}
}

//export gb2utf8
func gb2utf8(msg *C.char)(ret string){
    // gb2312转utf8
    cd, err := iconv.Open("utf-8", "gb2312");
    if err != nil{
    	log.Println("iconv open failed! ", err.Error());
    }
    defer cd.Close();
	
	ret = cd.ConvString(C.GoString(msg));
	return ret;
}

/*
go build -x -v -o gs
*/