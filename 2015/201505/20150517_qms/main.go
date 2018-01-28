package main

import(
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/braintree/manners"
	"fmt"
	"net/http"
)

 //http://localhost:8000/home?name=acb
func getting(ctx *gin.Context) {	
	name := ctx.DefaultQuery("name", "Guest");
	// name := ctx.Query("name");
	// name := ctx.Request.URL.Query().Get("name");
	fmt.Printf("Hello %s\n", name);
}

func posting(ctx *gin.Context){
	name := ctx.DefaultPostForm("name", "alert");
	fmt.Printf("Hello %s\n", name);
}

//localhost:8000/home/world
func api(ctx *gin.Context){
	name := ctx.Param("name");
	fmt.Printf("Hello %s\n", name);
}


func run(){
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	router := gin.Default();
	//创建不带中间件的路由：
	//r := gin.New()
	
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// r.Use(gin.Logger())
		
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	//r.Use(gin.Recovery())

	router.StaticFS("/m/", http.Dir("."));

	router.GET("/home", getting);
	router.GET("/api/:name", api);
	router.POST("/m/bus", posting); //http://localhost:8000/m/

	group := router.Group("/bus");
	group.GET("/bus", getting);
	group.POST("/bus", posting);

	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	
	manners.ListenAndServe(":8000", router);
}

func main(){
	run();
}