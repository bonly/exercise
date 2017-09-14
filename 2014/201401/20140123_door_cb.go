package main 

import (
// "fmt"
"log"
"net/http"
"io/ioutil"
)


func CallBack(rw http.ResponseWriter, qry *http.Request){
	log.Println("get a CallBack");
	body, err := ioutil.ReadAll(qry.Body);
	if err != nil{
		log.Println(err);
		return;
	}
	log.Println(string(body));
	return;
}

func main(){
	http.HandleFunc("/cb", CallBack);
	err := http.ListenAndServe(":9898", nil);
	if err != nil{
		log.Println(err);
		return;
	}

}

