package dal

import (
	"context"
	"douyinProject/video/cmd/rpc"
	"douyinProject/video/kitex_gen/user"
	"douyinProject/video/kitex_gen/video"
	"fmt"
	"time"
)

func CreateVideo(ctx context.Context, videoDB Video) error {

	tx := DB.Begin()

	res := tx.Create(&videoDB)

	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	rpcResp, err := rpc.UserClient.AddWorkNumMethod(ctx, &user.AddWorkNumReq{
		UserId: videoDB.AuthorId,
	})

	if err != nil {
		tx.Rollback()
		return err
	}
	code := rpcResp.StatusCode

	if code != 0 {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func NewFavoriteCount(vid, count int64) (int64, error) {
	var v Video

	res := DB.Where("id = ?", vid).Find(&v)
	if res.Error != nil {
		return 0, res.Error
	}
	v.FavoriteCount = v.FavoriteCount + count

	res = DB.Save(&v)
	if res.Error != nil {
		return 0, res.Error
	}
	return v.AuthorId, nil
}

func NewCommentCount(vid, count int64) error {
	var v Video

	res := DB.Where("id = ?", vid).Find(&v)
	if res.Error != nil {
		return res.Error
	}
	v.CommentCount = v.CommentCount + count

	res = DB.Save(&v)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetVideosByUserId(uid int64) ([]*video.Video, error) {

	var err error
	var videos = make([]*Video, 0)
	res := DB.Where("author_id = ?", uid).Find(&videos)

	var rvs = make([]*video.Video, 0)

	if res.Error != nil {
		return rvs, res.Error
	}

	rvs, err = PackVideos(videos)

	if err != nil {
		return nil, err
	}

	// 远程调用用户服务填充用户信息
	rpcResp, err := rpc.UserClient.InfoMethod(context.Background(), &user.InfoReq{UserId: uid})

	if err != nil {
		return nil, err
	}

	u := PackUser(rpcResp.User)

	// 回填结果集
	for _, v := range rvs {
		v.Author = u
	}
	return rvs, nil
}

func GetFeedVideos(t string, limit int) ([]*video.Video, int64, error) {

	var videos = make([]*Video, 0)

	res := DB.Where("created_at < ?", t).Limit(limit).Order("created_at DESC").Find(&videos)
	//res := DB.Where("created_at < ? and author_id != ?", t, uid).Limit(limit).Order("created_at DESC").Find(&videos)

	if res.Error != nil {
		return nil, time.Now().Unix(), fmt.Errorf("get video list from db fail")
	}
	if len(videos) == 0 {
		return nil, time.Now().Unix(), fmt.Errorf("get video list fail")
	}
	NextTime := videos[0].CreatedAt.Unix()

	resVL := make([]*video.Video, 0)

	for _, v := range videos {
		var video video.Video
		u, err := rpc.UserClient.InfoMethod(context.Background(), &user.InfoReq{UserId: v.AuthorId})

		if err != nil {
			return nil, time.Now().Unix(), fmt.Errorf("pack video list fail")
		}
		u1 := PackUser(u.User)
		video, err = PackVideo(v)
		video.Author = u1
		resVL = append(resVL, &video)
	}
	return resVL, NextTime, nil

}

func GetVideoList(vids []int64) ([]*video.Video, error) {
	var videos []Video

	res := DB.Where("id IN ?", vids).Find(&videos)

	if res.Error != nil {
		fmt.Println("error")
		return nil, res.Error
	}

	var videoList = make([]*video.Video, 0)

	for _, vid := range videos {
		vv, _ := PackVideo(&vid)
		uid := vid.AuthorId

		userInfo, err := rpc.UserClient.InfoMethod(context.Background(), &user.InfoReq{
			UserId: uid,
		})

		if err != nil {
			return nil, err
		}
		user := PackUser(userInfo.User)
		vv.Author = user
		videoList = append(videoList, &vv)
	}
	return videoList, nil

}
