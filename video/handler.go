package main

import (
	"bytes"
	"context"
	"douyinProject/video/cmd/consts"
	dal2 "douyinProject/video/cmd/dal"
	minio2 "douyinProject/video/cmd/minio"
	"douyinProject/video/cmd/rpc"
	"douyinProject/video/cmd/utils"
	"douyinProject/video/kitex_gen/action"
	"douyinProject/video/kitex_gen/relation"
	video "douyinProject/video/kitex_gen/video"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// FeedMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FeedMethod(ctx context.Context, request *video.FeedReq) (resp *video.FeedResp, err error) {

	videoNum := 1
	//uid := request.UserId
	timeStamp := request.LatestTime

	t := time.Unix(timeStamp/1000, timeStamp/10000000).Format("2006-01-02 15:04:05.000")

	videos, nextTime, err := dal2.GetFeedVideos(t, videoNum)
	//videos, nextTime, err := dal2.GetFeedVideos(t, videoNum, uid)
	if err != nil {
		return &video.FeedResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			VideoList:  nil,
			NextTime:   nextTime * 1000,
		}, err
	}

	for _, item := range videos {
		rpcResp1, err := rpc.RelationClient.IsFollowingMethod(ctx, &relation.IsFollowingReq{
			UserId:   request.UserId,
			ToUserId: item.Author.Id,
		})
		if err != nil {
			return &video.FeedResp{
				StatusCode: 1,
				StatusMsg:  err.Error(),
				VideoList:  nil,
				NextTime:   nextTime * 1000,
			}, err
		}
		rpcResp2, err := rpc.ActionClient.IsUserFavoriteVideoMethod(ctx, &action.IsUserFavoriteVideoReq{
			VideoId: item.Id,
			UserId:  request.UserId,
		})
		if err != nil {
			return &video.FeedResp{
				StatusCode: 1,
				StatusMsg:  err.Error(),
				VideoList:  nil,
				NextTime:   nextTime * 1000,
			}, err
		}
		// 当前登录用户对此视频是否点赞
		item.IsFavorite = rpcResp2.IsFavorite
		fType := rpcResp1.FollowingType
		if fType == 1 {
			item.Author.IsFollow = true
		} else {
			item.Author.IsFollow = false
		}
	}
	return &video.FeedResp{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
		NextTime:   nextTime * 1000,
	}, nil
}

// PublishActionMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishActionMethod(ctx context.Context, request *video.PublishActionReq) (resp *video.PublishActionResp, err error) {

	videoData := []byte(request.Data)

	// // 获取后缀
	//filetype := http.DetectContentType(videoData)
	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		msg := "upload fail"
		return &video.PublishActionResp{
			StatusCode: 1,
			StatusMsg:  &msg,
		}, err
	}
	fileName := u2.String() + "." + "mp4"
	// 开始上传文件
	err = minio2.UploadFile(ctx, consts.VideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		msg := "upload fail"
		return &video.PublishActionResp{
			StatusCode: 1,
			StatusMsg:  &msg,
		}, err
	}
	// 获取视频链接
	url, err := minio2.GetFileUrl(consts.VideoBucketName, fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]

	// 处理封面

	u3, err := uuid.NewV4()
	if err != nil {
		msg := "upload fail"
		return &video.PublishActionResp{
			StatusCode: 2,
			StatusMsg:  &msg,
		}, err
	}

	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	coverData, err := utils.ReadFrameAsJpeg(playUrl)
	if err != nil {
		msg := "get cover fail"
		return &video.PublishActionResp{
			StatusCode: 2,
			StatusMsg:  &msg,
		}, err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio2.UploadFile(ctx, consts.CoverBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		msg := "upload cover fail"
		return &video.PublishActionResp{
			StatusCode: 3,
			StatusMsg:  &msg,
		}, err
	}

	// 获取封面链接
	coverUrl, err := minio2.GetFileUrl(consts.CoverBucketName, coverPath, 0)
	if err != nil {
		msg := "get cover url fail"
		return &video.PublishActionResp{
			StatusCode: 4,
			StatusMsg:  &msg,
		}, err
	}

	CoverUrl := strings.Split(coverUrl.String(), "?")[0]

	videoDB := dal2.Video{
		AuthorId:      request.UserId,
		PlayUrl:       playUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         request.Title,
	}

	err = dal2.CreateVideo(ctx, videoDB)

	if err != nil {
		msg := "upload fail"
		return &video.PublishActionResp{
			StatusCode: 1,
			StatusMsg:  &msg,
		}, err
	}

	msg := "success"
	return &video.PublishActionResp{
		StatusCode: 0,
		StatusMsg:  &msg,
	}, err
}

// PublishListMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishListMethod(ctx context.Context, request *video.PublishListReq) (resp *video.PublishListResp, err error) {
	res, err := dal2.GetVideosByUserId(request.UserId)
	if err != nil {
		return nil, err
	}

	resp = new(video.PublishListResp)
	msg := "success"
	resp.StatusMsg = &msg
	resp.StatusCode = 0
	resp.VideoList = res
	return resp, nil
}

// FavoriteCountMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteCountMethod(ctx context.Context, request *video.FavoriteCountReq) (resp *video.FavoriteCountResp, err error) {
	// TODO: Your code here...
	vid := request.VideoId
	count := request.Count
	authorId, err := dal2.NewFavoriteCount(vid, count)

	if err != nil {
		return &video.FavoriteCountResp{
			StatusCode: 1,
			StatusMsg:  "update  video comment count fail",
			AuthorId:   authorId,
		}, err
	}
	return &video.FavoriteCountResp{
		StatusCode: 0,
		StatusMsg:  "success",
		AuthorId:   authorId,
	}, nil

}

// CommentCountMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentCountMethod(ctx context.Context, request *video.CommentCountReq) (resp *video.CommentCountResp, err error) {
	// TODO: Your code here...

	vid := request.VideoId
	count := request.Count
	err = dal2.NewCommentCount(vid, count)
	if err != nil {
		return &video.CommentCountResp{
			StatusCode: 1,
			StatusMsg:  "update video comment count fail",
		}, err
	}
	return &video.CommentCountResp{
		StatusCode: 0,
		StatusMsg:  "update video comment count success",
	}, nil
}

// VideoListMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoListMethod(ctx context.Context, request *video.VideoListReq) (resp *video.VideoListResp, err error) {
	vids := request.VideoIds
	res, _ := dal2.GetVideoList(vids)
	//var videos []dal2.Video
	//dal2.DB.Where("id IN ?", vids).Find(&videos)
	//fmt.Println(videos)
	//
	//fmt.Println("000000000000000000000000")
	if err != nil {
		return &video.VideoListResp{
			StatusCode: 1,
			StatusMsg:  "update video comment count fail",
			VideoList:  nil,
		}, err
	}
	return &video.VideoListResp{
		StatusCode: 0,
		StatusMsg:  "get user favorite vider list success",
		VideoList:  res,
	}, nil
}
