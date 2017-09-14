/*
auth: bonly
create: 2015.9.15
*/
package main 

import (
"fmt"
"net/http"
// "os"
"io/ioutil"
"path"
"flag"
)

var srv = flag.String("s", ":1997", "service address and port");

func main(){
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err);
        }
    }();

    flag.Parse();

	http.HandleFunc("/", Qry);
	err := http.ListenAndServe(*srv, nil);
	if err != nil{
		fmt.Println(err);
		return;
	}
}

func Qry(w http.ResponseWriter, r *http.Request){
	filename := path.Base(r.URL.Path) + ".json";
	fmt.Println("Qry: ", filename);

	dat, err := ioutil.ReadFile(filename);
	if err != nil{
		fmt.Println(err);
		return;
	}

	fmt.Println(string(dat));
	w.Write(dat);
}
