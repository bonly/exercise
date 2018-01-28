package main

import(
	"gopkg.in/gin-gonic/gin.v1"
	"time"
)

import (
	"net/http"	
	"github.com/fvbock/endless" // for endless 可无缝重启,windows不可用
	"github.com/braintree/manners" //可无缝重启，支持windows
)

func run_default(){
	router := gin.Default();
	router.StaticFS("/", http.Dir("."));
	router.Run(":8080");	
}

func run_http_default (){
	router := gin.Default();
	// http.Handle("/", http.FileServer(http.Dir(".")));
	router.StaticFS("/", http.Dir("."));
	http.ListenAndServe(":8080", router);
}

func run_manners(){
	router := gin.Default();
	router.StaticFS("/", http.Dir("."));
	manners.ListenAndServe(":8080", router);
}

func main(){
	// run_default();
	// run_http_default();
	// run_define();
	// run_endless();
	run_manners();
}

func run_define(){
	router := gin.Default();
	router.StaticFS("/", http.Dir("."));

	srv := &http.Server{
		Addr : ":8080",
		Handler : router,
		ReadTimeout : 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	};

	srv.ListenAndServe();
}

func run_endless(){ 
	router := gin.Default();
	router.StaticFS("/", http.Dir("."));
	endless.ListenAndServe(":8080", router);
}