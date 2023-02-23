package dal

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorId      int64  `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    bool   `gorm:"column:is_favorite;default:false"`
	Title         string `gorm:"column:title"`
}
