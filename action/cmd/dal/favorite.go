package dal

import (
	"fmt"
)

func FavoriteVideo(uid, vid int64, action int32) error {

	var ufv UserFavoriteVideo

	if action == 1 {
		ufv.VideoId = vid
		ufv.UserId = uid
		res := DB.Create(&ufv)
		return res.Error
	} else {
		res := DB.Where("video_id = ? and user_id = ?", vid, uid).Delete(&ufv)
		return res.Error
	}
}

func UserFavoriteVideoList(uid int64) ([]int64, error) {
	var videoIds = make([]int64, 0)
	var ufv = make([]*UserFavoriteVideo, 0)

	res := DB.Where("user_id = ?", uid).Find(&ufv)

	if res.Error != nil {
		return nil, res.Error
	}
	if len(ufv) == 0 {
		return nil, fmt.Errorf("user favorite videos are null")
	}
	for _, item := range ufv {
		vid := item.VideoId
		videoIds = append(videoIds, vid)

	}

	return videoIds, nil

}

func IsUserFavoriteVideo(vid, uid int64) (bool, error) {
	var ufv UserFavoriteVideo
	res := DB.Where("user_id = ?", uid).Where("video_id = ?", vid).Find(&ufv)
	return res.RowsAffected > 0, res.Error
}
