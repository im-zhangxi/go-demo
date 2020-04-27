package service

import (
	"github.com/jinzhu/gorm"
	"iddd.com/iddd-go/entity"
)

type roomService struct {
	db *gorm.DB
}

func (svc *roomService) Index() {
	svc.db.Create(&entity.Test1{
		Name: "name",
	})
}

func NewRoomService(db *gorm.DB) Room  {
	svc := new(roomService)
	svc.db = db

	return svc
}
