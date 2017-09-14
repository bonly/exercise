package main 

import (
// "net/http"
"github.com/gin-gonic/gin"
)

var router *gin.Engine

func  main(){
	gin.SetMode(gin.ReleaseMode);  //设置为生产模式

	router = gin.Default(); //创建

	router.LoadHTMLGlob("templates/*"); //创建所有模板

	//创建路由
	init_routes();

	router.Run(":8000");

}