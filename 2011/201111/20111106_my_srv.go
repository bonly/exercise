package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "sync"
)

type rest struct {
	Status string
	Code   int
	Msg    string
	Result string
}

func get(rw http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("imei");
	val := req.URL.Query().Get("simserial");
	fmt.Println("got: " + id + ": " + val);
	//rw.Write([]byte("got: " + id + ": " + val))
	var myrest rest;
	myrest.Status = "OK";
	myrest.Code = 200;
	myrest.Msg = "Success";
	myrest.Result = "";
	body, err := json.Marshal(myrest);
	if err != nil {
		panic(err.Error());
	}
    str := string(body);
    rw.Write([]byte(str));
}

// for real routing take a look at gorilla/mux package
func handler(rw http.ResponseWriter, req *http.Request) {
	get(rw, req)
	/*
	   switch req.Method {
	       case "GET":
	           get(rw, req);
	   }
	*/
}

func login(rw http.ResponseWriter, req *http.Request){
    fmt.Println("got name: " + req.FormValue("username"));
    fmt.Println("got passwd: " + req.FormValue("password"));
}

func main() {
    http.HandleFunc("/", handler);
    http.HandleFunc("/Login", login);
	err := http.ListenAndServe("localhost:80", nil);
	if err != nil {
		fmt.Println(err)
	}
}
