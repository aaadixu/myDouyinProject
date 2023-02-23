package main

import (
	"context"
	dal2 "douyinProject/user/cmd/dal"
	"douyinProject/user/cmd/utils"
	user "douyinProject/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

func (s *UserServiceImpl) FavoriteCountMethod(ctx context.Context, request *user.FavoriteCountReq) (r *user.FavoriteCountResp, err error) {

	uid := request.UserId
	count := request.Count

	err = dal2.FavoriteCount(uid, count)
	if err != nil {
		return &user.FavoriteCountResp{
			StatusCode: 1,
			StatusMsg:  "add favorite count fail",
		}, err
	}

	return &user.FavoriteCountResp{
		StatusCode: 0,
		StatusMsg:  "add favorite count scuuess",
	}, nil
}

func (s *UserServiceImpl) TotalFavoritedMethod(ctx context.Context, request *user.TotalFavoritedReq) (r *user.TotalFavoritedResp, err error) {
	uid := request.UserId

	count := request.Count

	err = dal2.TotalFavorited(uid, count)

	if err != nil {
		return &user.TotalFavoritedResp{
			StatusCode: 1,
			StatusMsg:  "add favorite count fail",
		}, err
	}

	return &user.TotalFavoritedResp{
		StatusCode: 0,
		StatusMsg:  "add favorite count scuuess",
	}, nil
}

// RegistMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegistMethod(ctx context.Context, request *user.RegistReq) (resp *user.RegistResp, err error) {
	username := request.Username
	password := request.Password

	var respponse user.RegistResp
	if dal2.ExistUserName(username) {
		var reqUser = dal2.User{
			Username: username,
			Password: utils.Md5(password),
		}
		us, err := dal2.CreateUser(reqUser)
		if err != nil {
			respponse.StatusCode = 1
			respponse.StatusMsg = "regist failed"
		} else {
			respponse.StatusCode = 0
			respponse.StatusMsg = "success"
			respponse.UserId = us.Id
		}

	} else {
		respponse.StatusCode = 1
		respponse.StatusMsg = "regist failed"
	}

	return &respponse, err

}

// LoginMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginMethod(ctx context.Context, request *user.LoginReq) (resp *user.LoginResp, err error) {
	var res user.LoginResp
	username := request.Username
	password := request.Password

	us, err := dal2.CheckUser(username, utils.Md5(password))

	if err != nil {
		res.StatusCode = 1
		res.StatusMsg = "login failed"
		res.UserId = 0
	} else {
		if us.Id == 0 {
			res.StatusCode = 1
			res.StatusMsg = "login failed"
			res.UserId = 0
		} else {
			res.StatusCode = 0
			res.UserId = us.Id
			res.StatusMsg = "login success"
		}
	}

	return &res, err
}

// InfoMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfoMethod(ctx context.Context, request *user.InfoReq) (resp *user.InfoResp, err error) {
	var res user.InfoResp
	uid := request.UserId
	us, err := dal2.UserInfo(uid) // 根据id查询用户信息
	if err != nil {
		res.StatusCode = 1
		res.StatusMsg = "get user info failed"
	} else {
		res.StatusCode = 0

		res.StatusMsg = "success"
		res.User = us
	}
	return &res, err
}

// AddWorkNumMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddWorkNumMethod(ctx context.Context, request *user.AddWorkNumReq) (resp *user.AddWorkNumResp, err error) {
	id := request.UserId
	err = dal2.UserAddWork(id)
	if err != nil {
		return &user.AddWorkNumResp{
			StatusCode: 1,
			StatusMsg:  "add user work num fail",
		}, err
	}
	return &user.AddWorkNumResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

// FollowCountMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowCountMethod(ctx context.Context, request *user.FollowCountReq) (resp *user.FollowCountResp, err error) {
	userId := request.UserId
	count := request.Count
	res := dal2.FollowCount(userId, count)
	if res != nil {
		return &user.FollowCountResp{
			StatusCode: 1,
			StatusMsg:  "update user follow count err",
		}, res
	}
	return &user.FollowCountResp{
		StatusCode: 0,
		StatusMsg:  "update user follow count success",
	}, nil
}

// FollowerCountMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowerCountMethod(ctx context.Context, request *user.FollowerCountReq) (resp *user.FollowerCountResp, err error) {
	userId := request.UserId
	count := request.Count
	res := dal2.FollowerCount(userId, count)
	if res != nil {
		return &user.FollowerCountResp{
			StatusCode: 1,
			StatusMsg:  "update user follower count err",
		}, res
	}
	return &user.FollowerCountResp{
		StatusCode: 0,
		StatusMsg:  "update user follower count success",
	}, nil
}

// InfosMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfosMethod(ctx context.Context, request *user.InfosReq) (resp *user.InfosResp, err error) {
	ids := request.UserIds

	res, err := dal2.UserInfos(ids)

	if err != nil {
		return &user.InfosResp{
			StatusCode: 1,
			StatusMsg:  "get user info list fail",
			Users:      nil,
		}, err
	}
	return &user.InfosResp{
		StatusCode: 0,
		StatusMsg:  "get user info list success",
		Users:      res,
	}, nil
}
