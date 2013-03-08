package main
import (
  "log"
  "fmt"
  "regexp"
  il "import_log" 
)

func main(){
	log.Println("begin");
	var fl il.MyFile;
	fl.Open ("gsrv.sxh-009.localdomain.wlmz.log.INFO.20130226-094214.12501");
	defer fl.Close();
	
	buf := make([]byte, 255);
  fl.ReadLineAt(buf, 14264);

  fmt.Printf("data:\n %s\n", buf);
  
  re, _ := regexp.Compile("{.*}");
  
  ch := re.Find(buf);
  fmt.Println("find: ", string(ch));
  
}
