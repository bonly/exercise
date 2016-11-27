package mt4

import (
"testing"
);

type St struct{
  Market string;
  Stock  string;
}

func spac(lst []St){
  for _, st := range lst {
    var dir string;
    if st.Market == "sz" {
      dir = "sz"
    }else{
      dir = "sh"
    }
    ifile := "/home/bonly/tdx/vipdoc/" + dir + "/lday/" + st.Market + st.Stock + ".day";
    ofile := "/home/bonly/mt4/history/ForexClub-MT4 Demo Server/" + 
             st.Market + st.Stock + "1440.hst";
    data2hst(ifile, ofile);
  }
}

func Test_file_audchf(ts *testing.T){
	read_hst("/home/bonly/mt4/history/ForexClub-MT4 Demo Server/AUDCHF1440.hst");
}


func Test_file_sz(ts *testing.T){
	read_hst("/home/bonly/htdzh/data/SZnse/Day/8199911440.hst");
}

func Test_150120(ts *testing.T){
	data2hst("/home/bonly/tdx/vipdoc/sz/lday/sz150120.day",
		"/home/bonly/mt4/history/ForexClub-MT4 Demo Server/sz1501201440.hst");
}

func Test_h510900(ts *testing.T){
	data2hst("/home/bonly/htdzh/data/SHase/Day/510900.day",
		"/home/bonly/mt4/history/ForexClub-MT4 Demo Server/sh5109001440.hst");
}


func Test_sz000100(ts *testing.T){
	data2hst("/home/bonly/htdzh/data/SZnse/Day/000100.day",
		"/home/bonly/mt4/history/ForexClub-MT4 Demo Server/sz0001001440.hst");
}

func Test_sh603288(ts *testing.T){
	data2hst("/home/bonly/htdzh/data/SHase/Day/603288.day",
		"/home/bonly/mt4/history/ForexClub-MT4 Demo Server/sh6032881440.hst");
}

func Test_hst_sz(ts *testing.T){
	build_hst_dir("sz","/home/bonly/htdzh/data/SZnse/Day",
		"/home/bonly/mt4/history/ForexClub-MT4 Demo Server/");
}

func Test_hst_sh(ts *testing.T){
	build_hst_dir("sh","/home/bonly/htdzh/data/SHase/Day",
		"/home/bonly/mt4/history/ForexClub-MT4 Demo Server/");
}

func Test_sp(ts *testing.T){
	lst := []St{
		{"sz","150120"},
		{"sh","510900"},
		{"sz","000100"},
		{"sh","603288"},
		{"sz","002024"},
		{"sz","002154"},
		{"sz","000539"},
		{"sz","000540"},
		{"sh","600894"},
	};
	spac(lst);
}