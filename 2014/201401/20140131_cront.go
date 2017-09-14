package main  

import (
"fmt"
"time"
cron "github.com/gorhill/cronexpr"
)

func main(){
	ac := cron.MustParse("* * * 23 12 * 2015").Next(time.Now());
	fmt.Println(ac);
	if ac.IsZero() == false{
	   fmt.Println("ok");
	}else{
	   fmt.Println("no");
	}
}