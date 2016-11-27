/*
The database header first . . . total 148 bytes

int	version;	// database version - 400	4 bytes
string	copyright[64];	// copyright info	64 bytes
string	symbol[12];	// symbol name	12 bytes
int	period;	// symbol timeframe	4 bytes
int	digits;	// the amount of digits after decimal point	4 bytes
datetime	timesign;	// timesign of the database creation	4 bytes
datetime	last_sync;	// the last synchronization time	4 bytes
int	unused[13];	// to be used in future	52 bytes

then the bars array (single-byte justification) . . . total 44 bytes

datetime	ctm;	// bar start time	4 bytes
double	open;	// open price	8 bytes
double	low;	// lowest price	8 bytes
double	high;	// highest price	8 bytes
double	close;	// close price	8 bytes
double	volume;	// tick count	8 bytes


.hst file format valid as of MT4 574 and later

The database header is the same . . . total 148 bytes

int			version;	// database version - 401	4 bytes
string		copyright[64];	// copyright info	64 bytes
string		symbol[12];	// symbol name	12 bytes
int			period;	// symbol timeframe	4 bytes
int			digits;	// the amount of digits after decimal point	4 bytes
datetime	timesign;	// timesign of the database creation	4 bytes
datetime	last_sync;	// the last synchronization time	4 bytes
int			unused[13];	// to be used in future	52 bytes

then the bars array (single-byte justification) . . . total 60 bytes

datetime	ctm;	// bar start time	8 bytes
double		open;	// open price	8 bytes
double		high;	// highest price	8 bytes
double		low;	// lowest price	8 bytes
double		close;	// close price	8 bytes
long		volume;	// tick count	8 bytes
int			spread;	// spread	4 bytes
long		real_volume;	// real volume	8 bytes


*/

package mt4

import (
	"fmt"
	"log"
	"os"
	"encoding/binary"
	"io"
  _"bufio"
	_"math"
	"time"
  "path/filepath"
  _"strconv"
  _"bytes"
  "strings"
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

func read_hst(file string){
	r, err := os.Open(file);
	if err != nil{
		log.Fatal("open feil failed: ", err);
	}
	defer r.Close();
	
	var head Head;
	if err = binary.Read(r, binary.LittleEndian, &head); err != nil{
		log.Panic("read head err: ", err);
	}
	fmt.Println("version: ",binary.LittleEndian.Uint32(head.Version[:]));
	fmt.Println(string(head.Copyright[:]));
	fmt.Println(string(head.Symbol[:]));
	fmt.Println("period: ", binary.LittleEndian.Uint32(head.Period[:]));
	fmt.Println("digits: ", binary.LittleEndian.Uint32(head.Digits[:]));
	fmt.Println("last_sync: ", binary.LittleEndian.Uint32(head.Last_sync[:]));
	t64 := (int64)(binary.LittleEndian.Uint32(head.Timesign[:]));
	fmt.Println("timesign: ", time.Unix(t64,0));
	// fmt.Println("timesign: ", time.Unix(int64(head.Timesign),0));
		
    for{
		var bar Bar;
		if err = binary.Read(r, binary.LittleEndian, &bar); err != nil{
			if err != io.EOF {log.Panic("get data err: ", err);}
			break;
		}
		fmt.Print("ctm: ", time.Unix(bar.Ctm, 0));
		fmt.Print("\topen: ", bar.Open);
		//fmt.Println(math.Float64frombits(binary.LittleEndian.Uint64(bar.Open[:])));
		fmt.Print("\thigh: ", bar.High);
		fmt.Print("\tlow: ", bar.Low);
		fmt.Print("\tclose: ", bar.Close);
		fmt.Print("\tvolume: ", bar.Volume);
		fmt.Print("\tspread: ", bar.Spread);
		fmt.Print("\treal_volume: ", bar.Real_volume);
		fmt.Println();
	}
}

/*
通达信数据
*/
/*
typedef struct mystructtag{    
   int date;    
   int open;    
   int high;   
   int low;    
   int close;    
   float amount;    
   int vol;     
   int reservation;
} StockData
*/
type Day struct{
  Date uint32;//date的格式：20070423
  Open uint32;//开盘价
  High uint32;//最高价
  Low uint32;//最低价
  Close uint32;//收盘价
  Amount float32;//成交金额
  Vol uint32;//成交数量
  Reservation uint32;//保留
};

func Stock(value []uint32, power []uint32) (res uint32){
	res = 0;
	for i := 1; i < len(value); i++{
		if value[i] >= value[i-1] { //价涨
		   if power[i] >= power[i-1] { //价涨量增 +2
		   		res += 2;
		   }else{  //价涨量缩 +1
			    res += 1;
		   }
		}else{ //价跌
			if power[i] >= power[i-1]{  //价跌量增 -2
			    res -= 2;
			}else{ //价跌量缩
				res -= 1;
			}
		}
	}
	return res;
}

func proc_file(file string, date uint32)(sum uint32, cnt uint32, price uint32, hight uint32){  
  r, err := os.Open(file);
  if err != nil {
    log.Fatal("open file failed: ",err);
  }
  defer r.Close();
  
  var value = make([]uint32,0,0);
  var power = make([]uint32,0,0);
  
  var day Day;
  for{
    if err = binary.Read(r, binary.LittleEndian, &day); err != nil{
      if err != io.EOF {log.Panic("get data err: ", err);}
      break;
    }
    
	  if date==0 || day.Date >= date{
//	    log.Println("data: ", day);
	    value = append(value, day.Close);
	    power = append(power, day.Vol);
      price = day.Close;
      if price > hight{
        hight = price;
      }
	  }
  }
//  log.Println("value: ", value);
//  log.Println("power: ", power);
  sum = Stock(value,power);
  cnt = uint32(len(value));
  return sum,cnt,price,hight;
}


func proc_dir(dir string, date uint32, min uint32){
	err := filepath.Walk(dir, 
	   func (path string, fl os.FileInfo, err error) error{
		   if (fl == nil){
			  log.Panic(err);
			  return err;
		   }
		   if (fl.IsDir()){
			  return nil; 
		   }
		   sum, cnt, price, hight := proc_file(path, date);
       if sum >= min {
		    log.Printf("[%s]cnt: %3d \tsum: %10d \tprice:%8.2f \thight:%.2f\n", 
                 path, cnt, sum, float32(price/1000.0), float32(hight/1000.0));
       }
		   return nil;
	   });
	if err != nil{
		log.Println("path err: ", err);
	}
}

func dump_dir(dir string, odir string){
  err := filepath.Walk(dir,
     func (path string, fl os.FileInfo, err error) error{
       if fl == nil{
         log.Panic(err);
         return err;
       }
       if (fl.IsDir()){
         return nil;
       }
       //ret, cnt := dump_file(path, odir+fl.Name()+".csv");
       _, _ = dump_file(path, odir+fl.Name()+".csv");
       //log.Printf("proc file ret %d, cnt: %d", ret, cnt);
       return nil;
     });
  if err != nil{
    log.Println("path err: ", err);
  }
}
func dump_file(ifile string, ofile string)(ret int, cnt int){
  //log.Println("begin process: ", ifile);
  if len(ofile)==0 {
    ofile = ifile + ".csv";
  }
  of, err := os.Create(ofile);
  if err != nil{
    log.Panic(err);
  }
  defer func(){
//    of.Flush();
    of.Sync();
    of.Close();
  }();

  r, err := os.Open(ifile);
  if err != nil {
    log.Fatal("open file failed: ",err);
  }
  defer r.Close();
  
  var day Day;
  for{
    if err = binary.Read(r, binary.LittleEndian, &day); err != nil{
      if err != io.EOF {
        log.Panic("get data err: ", err);
      }else{
//        log.Println("finish file: ", ofile);
      }
      break;
    }
    
    line := fmt.Sprintf("%d,%s,%d,%d,%d,%d,%d\n", 
                    day.Date,
                    "0:00",
                    day.Open,
                    day.High,
                    day.Low,
                    day.Close,
                    day.Vol);

    of.WriteString(line);
    cnt += 1;
  }
  return 0, cnt;
}

func build_hst_dir(pre string, dir string, odir string){
  err := filepath.Walk(dir,
     func (path string, fl os.FileInfo, err error) error{
       if fl == nil{
         log.Panic(err);
         return err;
       }
       if (fl.IsDir()){
         return nil;
       }
       of_name := odir + "/" +
         strings.TrimSuffix(filepath.Base(fl.Name()), filepath.Ext(fl.Name()))+
         "1440"+".hst"
       data2hst(path, of_name);
       //log.Printf("proc file ret %d, cnt: %d", ret, cnt);
       return nil;
     });
  if err != nil{
    log.Println("path err: ", err);
  }
}

func data2hst(ifile string, ofile string){
  if len(ifile) == 0{
    return;
  }
  if len(ofile)==0{
    ofile=filepath.Dir(ifile) + "/" +
          strings.TrimSuffix(filepath.Base(ifile),filepath.Ext(ifile)) + 
          "1440"+".hst";
  }
  of, err := os.Create(ofile);
  if err != nil{
    log.Panic(err);
  }
  defer func(){
    of.Sync();
    of.Close();
  }();
  
  r, err := os.Open(ifile);
  if err != nil{
    log.Fatal("open file failed: ", err);
  }
  defer r.Close();
  
  var head Head;
  binary.LittleEndian.PutUint32(head.Version[:], 401);
  copy(head.Copyright[:], "(C)opyright 2003, MetaQuotes Software Corp.");
  copy(head.Symbol[:], strings.TrimSuffix(filepath.Base(ifile),filepath.Ext(ifile)));
  binary.LittleEndian.PutUint32(head.Period[:], 1440);
  binary.LittleEndian.PutUint32(head.Digits[:], 5);
  binary.LittleEndian.PutUint32(head.Timesign[:], (uint32)(time.Now().Unix()));
  
  if err = binary.Write(of, binary.LittleEndian, &head); err!=nil{
    log.Panic("write head err: ", err);
  }
  
  var day Day;
  for {
    if err = binary.Read(r, binary.LittleEndian, &day); err!=nil{
      if err != io.EOF{
        log.Panic("get data err: ",err);
      }else{
        //finish
      }
      break;
    }
    
    // fmt.Println(day);
    var bar Bar;
    str := fmt.Sprintf("%d 00:00:00", day.Date);
    tm,err := time.Parse("20060102 15:04:05", str); //01/02 03:04:05PM '06 -0700
    if err != nil{
      log.Println("time err: ", err);
    }
    bar.Ctm = tm.Unix();
    bar.Open = float64(day.Open) / 100.0;
    bar.High = float64(day.High) / 100.0;
    bar.Low = float64(day.Low) / 100.0;
    bar.Close = float64(day.Close) / 100.0;
    // bar.Open = float64(day.Open);
    // bar.High = float64(day.High);
    // bar.Low = float64(day.Low);
    // bar.Close = float64(day.Close);    
    bar.Volume = int64(day.Vol);
    // // bar.Spread
    // // bar.Real_volume
    
    if err = binary.Write(of, binary.LittleEndian, &bar); err!=nil{
      log.Panic("write bar err: ", err);
      break;
    }
  }
}

