package main 

import (
"net/http"
"github.com/gin-gonic/gin"
)

var router *gin.Engine

func  main(){
	gin.SetMode(gin.ReleaseMode);  //设置为生产模式

	router = gin.Default(); //创建

	router.LoadHTMLGlob("templates/*"); //创建所有模板

	//创建路由
	router.GET("/", func(c *gin.Context){
		c.HTML(
			http.StatusOK, //设置200（OK)状态响应值
			"index.html", //回应内容
			gin.H{
				"title": "Home Page",
			},
		);
	});

	router.Run();

}