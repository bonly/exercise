/*
auth: bonly
create: 2015.9.15
*/
package main 

import (
"fmt"
"net/http"
"os"
"io/ioutil"
)

func main(){
	if len(os.Args) != 3{
		fmt.Println("useage: ", os.Args[0], " ip:port file");
		return;
	}
	http.HandleFunc("/api/order/queryOrder.do", Qry);
	err := http.ListenAndServe(os.Args[1], nil);
	if err != nil{
		fmt.Println(err);
		return;
	}
}

func Qry(w http.ResponseWriter, r *http.Request){
	dat, err := ioutil.ReadFile(os.Args[2]);
	if err != nil{
		fmt.Println(err);
		return;
	}

	fmt.Println(string(dat));
	w.Write(dat);
}