package main

/*
#cgo CPPFLAGS: -g
#cgo LDFLAGS: -lbonly -lthostmduserapi -lthosttraderapi -L.
#include "acc.h"
#include "ThostFtdcUserApiStruct.h"
*/
import "C"
import (
	"fmt"
	"os"
	"log"
	"encoding/binary"
	// "path/filepath"
	// "strings"
	"time"
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

//export Hello
func Hello(){
	fmt.Println("hello");
}


func main(){
	fmt.Println("begin");
	C.cnt();
	fmt.Println("end");
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

