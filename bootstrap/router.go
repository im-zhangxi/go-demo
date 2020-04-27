package bootstrap

import (

)

func initRouter() {
	appServer.GET("/room", roomController.Index)
}
