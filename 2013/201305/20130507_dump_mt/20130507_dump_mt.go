package stock

import (
	"log"
	"os"
	"io"
	"encoding/binary"
	"path/filepath"
//  "bufio"
//  "io/ioutil"
  "fmt"
)

/*
type Day struct{
  unsigned long date;//date的格式：20070423
  unsigned long open;//开盘价
  unsigned long high;//最高价
  unsigned long low;//最低价
  unsigned long close;//收盘价
  unsigned long moneysum;//成交金额
  unsigned long turnover;//成交数量
  char unused[12];//保留
};
*/
type Day struct{
  Date uint32;//date的格式：20070423
  Open uint32;//开盘价
  High uint32;//最高价
  Low uint32;//最低价
  Close uint32;//收盘价
  Moneysum uint32;//成交金额
  Turnover uint32;//成交数量
  Unused [12]byte;//保留
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
	    power = append(power, day.Turnover);
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
                    day.Turnover);

    of.WriteString(line);
    cnt += 1;
  }
  return 0, cnt;
}
