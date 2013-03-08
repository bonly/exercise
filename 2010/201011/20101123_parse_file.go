package main
import (
  "log"
  "fmt"
  "regexp"
  "io"
  "strings"
  "flag"
  il "import_log" 
)

func main(){
	log.Println("begin");
	file_name := flag.String("c","gsrv.sxh-009.localdomain.wlmz.log.INFO.20130226-094214.12501","gsrv log file's name");
	flag.Parse();
	
	var fl il.MyFile;
	fl.Open (*file_name);
	//fl.Open ("20101124_File_Manager.go");
	defer fl.Close();
	
	var db il.MyDb;
	db.Open("127.0.0.1");
	defer db.Close();
	
	for {
		if str, err := fl.ReadLineAt(fl.Pos); err == nil {
		  //log.Printf("%s\n", str);
  
		  re, _ := regexp.Compile("{.*}");
		  
		  by := []byte(strings.Join(str,""));
		  ch := re.Find(by);
		  if ch != nil{
		    //fmt.Println(string(ch));
		    db.Write_log(ch);
		    
		  }

		}	else if err == io.EOF {
	     err = nil;
	     return;
	  }else{
			fmt.Println(err);
			break;
		}
	}
	
	log.Println("end");
  
}
