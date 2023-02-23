package dal

import (
	"douyinProject/chat/cmd/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(consts.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&ChatRecord{})
	if err != nil {
		panic("failed to create table")
	}

}
