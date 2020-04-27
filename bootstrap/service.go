package bootstrap

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"iddd.com/iddd-go/service"
	"log"
)

var (
	db          *gorm.DB
	roomService service.Room
)

func getDB() *gorm.DB {
	connect, err := gorm.Open("mysql", appConfig.DB.URL)
	if err != nil {
		log.Fatal(err)
	}
	db = connect
	db.DB().SetMaxIdleConns(10) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	db.DB().SetMaxOpenConns(50) //设置数据库连接池最大连接数

	return db
}

func initService() {
	db = getDB()
	roomService = service.NewRoomService(db)
}
