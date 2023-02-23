package dal

import (
	"douyinProject/video/kitex_gen/user"
	"douyinProject/video/kitex_gen/video"
)

func PackVideo(v *Video) (video.Video, error) {
	var rv video.Video
	rv.CommentCount = v.CommentCount

	rv.Id = int64(v.ID)
	rv.FavoriteCount = v.FavoriteCount
	rv.Title = v.Title
	rv.PlayUrl = v.PlayUrl
	rv.CoverUrl = v.CoverUrl
	rv.IsFavorite = v.IsFavorite
	return rv, nil
}

func PackVideos(vs []*Video) ([]*video.Video, error) {
	var rvs = make([]*video.Video, 0)

	for _, v := range vs {
		v1, _ := PackVideo(v)
		rvs = append(rvs, &v1)
	}
	return rvs, nil
}

func PackUser(u *user.User) *video.User {
	return &video.User{
		Id:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
}
