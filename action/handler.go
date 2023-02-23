package main

import (
	"context"
	"douyinProject/action/cmd/dal"
	"douyinProject/action/cmd/rpc"
	action "douyinProject/action/kitex_gen/action"
	"douyinProject/action/kitex_gen/user"
	"douyinProject/action/kitex_gen/video"
	"fmt"
)

// ActionServiceImpl implements the last service interface defined in the IDL.
type ActionServiceImpl struct{}

// FavoriteActionMethod implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) FavoriteActionMethod(ctx context.Context, request *action.FavoriteActionReq) (resp *action.FavoriteActionResp, err error) {

	actionType := request.ActionType

	vid := request.VideoId

	var count int8

	if actionType == 1 {
		count = 1
	}
	if actionType == 2 {
		count = -1
	}

	// 点赞操作

	// 1 点赞用户的 favorite_count + 1 （remote）
	rpcResp, err := rpc.UserClient.FavoriteCountMethod(ctx, &user.FavoriteCountReq{
		UserId: request.UserId,
		Count:  count,
	})
	if err != nil {
		return &action.FavoriteActionResp{
			StatusCode: rpcResp.StatusCode,
			StatusMsg:  rpcResp.StatusMsg,
		}, err
	}

	// 3 视频获赞 favorite_count + 1 （remote） 返回authorId

	rpcRespV, err := rpc.VideoClient.FavoriteCountMethod(ctx, &video.FavoriteCountReq{
		VideoId: request.VideoId,
		Count:   int64(count),
	})
	if err != nil {
		return &action.FavoriteActionResp{
			StatusCode: rpcRespV.StatusCode,
			StatusMsg:  rpcRespV.StatusMsg,
		}, err
	}
	authorId := rpcRespV.AuthorId
	// 2 被点赞视频的作者 total_favorited + 1 （remote）
	rpcRespU, err := rpc.UserClient.TotalFavoritedMethod(ctx, &user.TotalFavoritedReq{
		UserId: authorId,
		Count:  count,
	})

	if err != nil {
		return &action.FavoriteActionResp{
			StatusCode: rpcRespU.StatusCode,
			StatusMsg:  rpcRespU.StatusMsg,
		}, err
	}

	// 4 用户喜欢列表新增一条记录 (local)

	err = dal.FavoriteVideo(request.UserId, vid, actionType)
	if err != nil {
		return &action.FavoriteActionResp{
			StatusCode: 3,
			StatusMsg:  "update user favorite video table fail",
		}, err
	}
	return &action.FavoriteActionResp{
		StatusCode: 0,
		StatusMsg:  "update user favorite video table success",
	}, nil
}

// FavoriteListMethod implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) FavoriteListMethod(ctx context.Context, request *action.FavoriteListReq) (resp *action.FavoriteListResp, err error) {

	uid := request.UserId

	vids, err := dal.UserFavoriteVideoList(uid)
	if err != nil {
		return &action.FavoriteListResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		}, err
	}

	videoList, err := rpc.VideoClient.VideoListMethod(ctx, &video.VideoListReq{VideoIds: vids})
	if err != nil {
		return &action.FavoriteListResp{
			StatusCode: 2,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		}, err
	}

	vdList, err := dal.PackVideos(videoList.VideoList)

	return &action.FavoriteListResp{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  vdList,
	}, nil
}

// CommentActionMethod implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) CommentActionMethod(ctx context.Context, request *action.CommentActionReq) (resp *action.CommentActionResp, err error) {

	// 如果是发布评论
	var count int64
	if request.ActionType == 1 {
		count = 1
		// 视频评论数变化
		rpcResp, err := rpc.VideoClient.CommentCountMethod(ctx, &video.CommentCountReq{
			VideoId: request.VideoId,
			Count:   count,
		})
		if err != nil {
			return &action.CommentActionResp{
				StatusCode: 3,
				StatusMsg:  "update video comment count fail",
				Comment:    nil,
			}, nil
		}
		// 添加到数据库
		commnet, err := dal.AddComment(request.UserId, request.VideoId, request.CommentText)
		if err != nil {
			return &action.CommentActionResp{
				StatusCode: 3,
				StatusMsg:  "add comment fail",
				Comment:    commnet,
			}, err
		}
		rpcUser, err := rpc.UserClient.InfoMethod(ctx, &user.InfoReq{UserId: request.UserId})
		if err != nil {
			return &action.CommentActionResp{
				StatusCode: rpcResp.StatusCode,
				StatusMsg:  rpcResp.StatusMsg,
				Comment:    nil,
			}, nil
		}

		user, err := dal.PackUser(rpcUser.User)
		commnet.User = user
		return &action.CommentActionResp{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    commnet,
		}, nil
	} else { // 删除评论
		// 首先判断要删除的评论的作者id和当前用户id是否一致
		userId := request.UserId
		commnentId := request.CommentId
		commAuthorId, err := dal.GetCommentAuthorId(*commnentId)
		if err != nil {
			return &action.CommentActionResp{
				StatusCode: 100,
				StatusMsg:  "get  comment  atuthor id  fail",
				Comment:    nil,
			}, err
		}

		if userId != commAuthorId {
			return &action.CommentActionResp{
				StatusCode: 100,
				StatusMsg:  "delete  comment  fail",
				Comment:    nil,
			}, fmt.Errorf("用户id与评论作者id不一致，不能删除评论")
		}

		count = -1
		// 视频评论数变化
		rpcResp, err := rpc.VideoClient.CommentCountMethod(ctx, &video.CommentCountReq{
			VideoId: request.VideoId,
			Count:   count,
		})

		if err != nil {
			return &action.CommentActionResp{
				StatusCode: rpcResp.StatusCode,
				StatusMsg:  "update video comment count fail",
				Comment:    nil,
			}, err
		}
		// 删除评论
		res := dal.DeleteComment(request.VideoId, *request.CommentId)
		if res != nil {
			return &action.CommentActionResp{
				StatusCode: rpcResp.StatusCode,
				StatusMsg:  "delete video comment fail",
				Comment:    nil,
			}, err
		}
		return &action.CommentActionResp{
			StatusCode: 0,
			StatusMsg:  "delete video comment success",
			Comment:    nil,
		}, err
	}

}

// CommentListMethod implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) CommentListMethod(ctx context.Context, request *action.CommentListReq) (resp *action.CommentListResp, err error) {
	videoId := request.VideoId

	res, err := dal.GetCommentList(videoId)

	if err != nil {
		return &action.CommentListResp{
			StatusCode:  1,
			StatusMsg:   "get comment list fail",
			CommentList: nil,
		}, err
	}

	return &action.CommentListResp{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: res,
	}, nil
}

// IsUserFavoriteVideoMethod implements the ActionServiceImpl interface.
func (s *ActionServiceImpl) IsUserFavoriteVideoMethod(ctx context.Context, request *action.IsUserFavoriteVideoReq) (resp *action.IsUserFavoriteVideoResp, err error) {
	// TODO: Your code here...
	uid := request.UserId
	vid := request.VideoId
	res, err := dal.IsUserFavoriteVideo(vid, uid)
	if err != nil {
		return &action.IsUserFavoriteVideoResp{
			StatusCode: 1,
			StatusMsg:  "get user favorite status for video fail",
			IsFavorite: false,
		}, err
	}
	return &action.IsUserFavoriteVideoResp{
		StatusCode: 0,
		StatusMsg:  "get user favorite status for video success",
		IsFavorite: res,
	}, nil
}
