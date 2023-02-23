// Code generated by hertz generator.

package http

import (
	"context"
	http "douyinProject/http/biz/model/http"
	jwt "douyinProject/http/biz/mw"
	"douyinProject/http/biz/rpc"
	"douyinProject/http/kitex_gen/httprpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommentActionMethod .
// @router /douyin/comment/action/ [POST]
func CommentActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req http.CommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(http.CommentActionResp)

	res, err := jwt.AuthToken(ctx, c)
	if err != nil || res == -1 { // token校验失败
		msg := "auth fail"
		resp.StatusCode = 2
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	rpcResp, err := rpc.ActionClient.CommentActionMethod(ctx, &httprpc.CommentActionReq{
		UserId:      res,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	})
	resp.StatusCode = rpcResp.StatusCode
	resp.StatusMsg = &rpcResp.StatusMsg

	comm, err := rpc.PackComment(rpcResp.Comment)

	resp.Comment = comm
	c.JSON(consts.StatusOK, resp)

}

// CommentListMethod .
// @router /douyin/comment/list/ [GET]
func CommentListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req http.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(http.CommentListResp)

	rpcResp, err := rpc.ActionClient.CommentListMethod(ctx, &httprpc.CommentListReq{VideoId: req.VideoID})

	if err != nil {
		msg := "get comment list fail"
		c.JSON(consts.StatusOK, &http.CommentListResp{
			StatusCode:  1,
			StatusMsg:   &msg,
			CommentList: nil,
		})
		return
	}

	resp.StatusCode = rpcResp.StatusCode
	resp.StatusMsg = &rpcResp.StatusMsg

	comments, err := rpc.PackComments(rpcResp.CommentList)
	if err != nil {
		msg := "pack comment list fail"
		c.JSON(consts.StatusOK, &http.CommentListResp{
			StatusCode:  1,
			StatusMsg:   &msg,
			CommentList: nil,
		})
		return
	}
	resp.CommentList = comments

	c.JSON(consts.StatusOK, resp)
}

// FavoriteActionMethod .
// @router /douyin/favorite/action/ [POST]
func FavoriteActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req http.FavoriteActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(http.FavoriteActionResp)

	res, err := jwt.AuthToken(ctx, c)
	if err != nil || res == -1 { // token校验失败
		msg := "auth fail"
		resp.StatusCode = 2
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	rpcResp, err := rpc.ActionClient.FavoriteActionMethod(ctx, &httprpc.FavoriteActionReq{
		UserId:     res,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})

	if err != nil {
		resp.StatusCode = 1
		msg := "点赞失败"
		resp.StatusMsg = &msg
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.StatusCode = rpcResp.StatusCode
	resp.StatusMsg = &rpcResp.StatusMsg
	c.JSON(consts.StatusOK, resp)
}

// FavoriteListMethod .
// @router /douyin/favorite/list/ [GET]
func FavoriteListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req http.FavoriteListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(http.FavoriteListResp)

	res, err := jwt.AuthToken(ctx, c)
	if err != nil || res == -1 { // token校验失败
		msg := "auth fail"
		resp.StatusCode = 2
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	rpcResp, err := rpc.ActionClient.FavoriteListMethod(ctx, &httprpc.FavoriteListReq{UserId: req.UserID})

	if err != nil {
		msg := "get list info err"
		c.JSON(consts.StatusBadRequest, &http.FavoriteListResp{
			StatusCode: 1,
			StatusMsg:  &msg,
			VideoList:  nil,
		})
		return
	}

	videos, err := rpc.PackVideos(rpcResp.VideoList)
	if err != nil {
		msg := "pack list info err"
		c.JSON(consts.StatusBadRequest, &http.FavoriteListResp{
			StatusCode: 2,
			StatusMsg:  &msg,
			VideoList:  nil,
		})
		return
	}
	resp.VideoList = videos
	c.JSON(consts.StatusOK, resp)
}
