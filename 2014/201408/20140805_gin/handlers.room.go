package main 

import (
"net/http"
"github.com/gin-gonic/gin"
"strconv"
"log"
// "fmt"
)

func showIndexPage(cx *gin.Context){
	rooms := getAllRoom();

	cx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title" : "Xbed",
			"data": rooms,
		},
	);
}

func getRoom(cx *gin.Context){
	if RoomID, err := strconv.Atoi(cx.Param("room")); err != nil{
		log.Println("view room: ", RoomID);
		if room, err := getRoomByID(RoomID); err == nil{
			log.Println("found room: ", RoomID);
			cx.HTML(
				http.StatusOK,
				"room.html",
				gin.H{
					"Address": room.Address,
					"data": room,
				},
			);
		}else{
			log.Println(err);
			cx.AbortWithError(http.StatusNotFound, err);
		}
	}else{
		cx.AbortWithStatus(http.StatusNotFound);
	}
}