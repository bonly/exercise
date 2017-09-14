package main 

func init_routes() {
	router.GET("/", showIndexPage);

	view := router.Group("/view");
	view.GET("/room/:roomid", getRoom);
}