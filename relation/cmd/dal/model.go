package dal

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowId   int64 `gorm:"column:follow_id"`
	FollowerId int64 `gorm:"column:follower_id"`
}

type Friend struct {
	gorm.Model
	User1Id int64 `gorm:"column:user1_id"`
	User2Id int64 `gorm:"column:user2_id"`
}
