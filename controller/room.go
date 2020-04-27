package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iddd.com/iddd-go/service"
	"net/http"
)

type Room interface {
	Index(*gin.Context)
}

type roomController struct {
	roomService service.Room
}

func (ctrl *roomController) Index(ctx *gin.Context) {
	fmt.Println("xxxxxxxxxxxxxxx")
	ctrl.roomService.Index()

	ctx.JSON(http.StatusOK, gin.H{"message": "message"})
	return
}

func NewRoomController(roomService service.Room) Room {
	ctrl := new(roomController)
	ctrl.roomService = roomService

	return ctrl
}
