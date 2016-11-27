package stock

import (
"testing"
);


func Test_深华新(ts *testing.T){
  var value = []uint32{814,895,985,1084,1192,1311};
  var power = []uint32{2385,722,517,1094,1389,1319};
  res :=  Stock(value, power);
  ts.Log("深华新6天总分为：", res);
}


func Test_汉麻产业(ts *testing.T){
  var value = []uint32{828,830,863,859,945,1040};
  var power = []uint32{21025,22602,67180,65209,757,2290};
  res:=  Stock(value, power);
  ts.Log("汉麻产业6天总分为：", res);
}

func Test_file(ts *testing.T){
  sum, cnt, _,_ := proc_file("/home/bonly/htdzh/data/SZnse/Day/399666.day",0);
  ts.Log("统计数： ",cnt,"\t得分: ", sum, "\n");
}

func Test_dir(ts *testing.T){
  proc_dir("/home/bonly/htdzh/data/SZnse/Day",20150401, 30);
}

func Test_sh(ts *testing.T){
  proc_dir("/home/bonly/htdzh/data/SHase/Day",20150401, 30);
}

func Test_819991(ts *testing.T){
  sum, cnt, _,_ := proc_file("/home/bonly/htdzh/data/SZnse/Day/819991.day",0);
  ts.Log("统计数： ",cnt,"\t得分: ", sum, "\n");
}

func Test_399003(ts *testing.T){
  sum, cnt, _,_ := proc_file("/home/bonly/htdzh/data/SZnse/Day/399003.day",0);
  ts.Log("统计数： ",cnt,"\t得分: ", sum, "\n");
}

func Test_600839(ts *testing.T){
  sum, cnt, price, hight := proc_file("/home/bonly/htdzh/data/SHase/Day/600839.day",0);
  ts.Log("统计数： ",cnt,"\t得分: ", sum, "\t当前价: ", price, "\t最高价: ", hight);
}

func Test_200058(ts *testing.T){
  sum, cnt, price, hight := proc_file("/home/bonly/htdzh/data/SZnse/Day/200058.day",0);
  ts.Log("统计数： ",cnt,"\t得分: ", sum, "\t当前价: ", price, "\t最高价: ", hight);
}

func Test_dump_200058(ts *testing.T){
  ret, cnt := dump_file("/home/bonly/htdzh/data/SZnse/Day/200058.day","");
  ts.Log("ret: ",ret ," cnt: ", cnt);
}

func Test_dump_000001(ts *testing.T){
  ret, cnt := dump_file("/home/bonly/htdzh/data/SZnse/Day/000001.day","");
  ts.Log("ret: ",ret ," cnt: ", cnt);
}

func Test_dump_sz(ts *testing.T){
  dump_dir("/home/bonly/htdzh/data/SZnse/Day", "/home/bonly/stock_data_sz/");
}

func Test_dump_sh(ts *testing.T){
  dump_dir("/home/bonly/htdzh/data/SHase/Day", "/home/bonly/stock_data_sh/");
}
/*
go test -v -test.run dir 2> res.log
*/
