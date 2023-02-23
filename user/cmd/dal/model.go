package dal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `gorm:"column:user_name"`
	Password        string `gorm:"column:password"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	IsFollow        bool   `gorm:"column:is_follow"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFavorited  int64  `gorm:"column:total_favorited"`
	WorkCount       int64  `gorm:"column:work_count"`
	FavoriteCount   int64  `gorm:"column:favorite_count"`
}
