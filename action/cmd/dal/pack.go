package dal

import (
	"douyinProject/action/kitex_gen/action"
	"douyinProject/action/kitex_gen/user"
	"douyinProject/action/kitex_gen/video"
)

func PackVideo(v *video.Video) (action.Video, error) {
	var rv action.Video
	rv.CommentCount = v.CommentCount

	rv.Id = v.Id
	rv.FavoriteCount = v.FavoriteCount
	rv.Title = v.Title
	rv.PlayUrl = v.PlayUrl
	rv.CoverUrl = v.CoverUrl
	rv.IsFavorite = v.IsFavorite
	return rv, nil
}

func PackVideos(vs []*video.Video) ([]*action.Video, error) {
	res := make([]*action.Video, 0)
	for _, item := range vs {
		vd, _ := PackVideo(item)
		res = append(res, &vd)
	}
	return res, nil
}

func PackUser(user *user.User) (*action.User, error) {
	return &action.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}, nil
}

func PackComment(comm Comment) (*action.Comment, error) {
	return &action.Comment{
		Id:         int64(comm.ID),
		User:       nil,
		Content:    comm.Content,
		CreateDate: comm.CreatedAt.Format("01-02"),
	}, nil
}

func PackComments(comms []Comment) ([]*action.Comment, error) {
	res := make([]*action.Comment, 0)

	for _, c := range comms {
		cc, _ := PackComment(c)
		res = append(res, cc)
	}
	return res, nil
}
