package bootstrap

import (
	"iddd.com/iddd-go/controller"
)

var (
	roomController controller.Room
)

func initController() {
	roomController = controller.NewRoomController(roomService)
}
