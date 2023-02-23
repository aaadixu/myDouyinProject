package dal

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	VideoId int64
	UserId  int64
	Content string `gorm:"column:content"`
}

type UserFavoriteVideo struct {
	gorm.Model
	UserId  int64
	VideoId int64
}
