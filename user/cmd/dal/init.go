package dal

import (
	"douyinProject/user/cmd/consts"
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

	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic("failed to create table")
	}

}
