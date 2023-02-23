package dal

import "gorm.io/gorm"

type ChatRecord struct {
	gorm.Model
	FromUserId int64
	ToUserId   int64
	Content    string
	CreateTime string
}
