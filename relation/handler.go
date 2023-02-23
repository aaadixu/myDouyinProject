package main

import (
	"context"
	"douyinProject/relation/cmd/dal"
	"douyinProject/relation/cmd/rpc"
	relation "douyinProject/relation/kitex_gen/relation"
	"douyinProject/relation/kitex_gen/user"
	"fmt"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationActionMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationActionMethod(ctx context.Context, request *relation.RelationActionReq) (resp *relation.RelationActionResp, err error) {

	userId := request.UserId     // 我是粉丝follower
	toUserId := request.ToUserId // 对方是follow

	var count int64

	if request.ActionType == 1 { //关注操作
		count = 1
		// follow添加一条记录
		res, _ := dal.IsFollow(toUserId, userId)
		if !res { // 判断是否已经添加记录，避免重复关注
			err := dal.AddFollow(toUserId, userId)
			if err != nil {
				return &relation.RelationActionResp{
					StatusCode: 1,
					StatusMsg:  "add relation action fail",
				}, err
			}
		}

		// 我自己的follow加1
		rpcResp, err := rpc.UserClient.FollowCountMethod(ctx, &user.FollowCountReq{
			UserId: userId,
			Count:  count,
		})
		if err != nil {
			return &relation.RelationActionResp{
				StatusCode: rpcResp.StatusCode,
				StatusMsg:  rpcResp.StatusMsg,
			}, err
		}

		// 被关注者（对方）的follower加1
		rpcResp1, err := rpc.UserClient.FollowerCountMethod(ctx, &user.FollowerCountReq{
			UserId: toUserId,
			Count:  count,
		})
		if err != nil {
			return &relation.RelationActionResp{
				StatusCode: rpcResp1.StatusCode,
				StatusMsg:  rpcResp1.StatusMsg,
			}, err
		}
		// 然后判断两个人是否相互关注,只需要判断对方是否关注我即可
		res, err = dal.IsFollow(userId, toUserId)
		if err != nil {
			return &relation.RelationActionResp{
				StatusCode: 2,
				StatusMsg:  "get mutual relation fail",
			}, err
		}
		if res { // 为true说明当前二者已经相互关注，则二人互为朋友
			res := dal.AddFriend(userId, toUserId)
			if res != nil {
				return &relation.RelationActionResp{
					StatusCode: 3,
					StatusMsg:  "add friend  relation fail",
				}, err
			}
		}
		return &relation.RelationActionResp{
			StatusCode: 0,
			StatusMsg:  "success",
		}, nil

	} else { // 取消关注
		count = -1
		// follow 删除关注记录
		res := dal.DeleteFollow(toUserId, userId)
		if res != nil {
			return &relation.RelationActionResp{
				StatusCode: 4,
				StatusMsg:  "delete relation fail",
			}, err
		}

		// 我的关注减1（follow）
		rpcResp, err := rpc.UserClient.FollowCountMethod(ctx, &user.FollowCountReq{
			UserId: userId,
			Count:  count,
		})
		if err != nil {
			return &relation.RelationActionResp{
				StatusCode: rpcResp.StatusCode,
				StatusMsg:  rpcResp.StatusMsg,
			}, err
		}
		// 对方的粉丝数减1（follower）
		rpcResp1, err := rpc.UserClient.FollowerCountMethod(ctx, &user.FollowerCountReq{
			UserId: toUserId,
			Count:  count,
		})
		if err != nil {
			return &relation.RelationActionResp{
				StatusCode: rpcResp1.StatusCode,
				StatusMsg:  rpcResp1.StatusMsg,
			}, err
		}
		// 判断是否存在朋友关系，有则删除
		ship := dal.ExistedFriendship(userId, toUserId)

		if ship { // 存在朋友关系
			res = dal.DeleteFriend(userId, toUserId)
			if res != nil {
				return &relation.RelationActionResp{
					StatusCode: 10,
					StatusMsg:  "delete ffriend relation err",
				}, res
			}
		}
		return &relation.RelationActionResp{
			StatusCode: 0,
			StatusMsg:  "success",
		}, nil
	}
}

// FollowListMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowListMethod(ctx context.Context, request *relation.FollowListReq) (resp *relation.FollowListResp, err error) {
	userId := request.UserId
	followIds, err := dal.GetFollowList(userId)
	if err != nil {
		return &relation.FollowListResp{
			StatusCode: 1,
			StatusMsg:  "get follow ids fail",
			UserList:   nil,
		}, nil
	}

	fmt.Println("666666666666666666666666666666666", followIds)

	usersResp, err := rpc.UserClient.InfosMethod(ctx, &user.InfosReq{UserIds: followIds})
	fmt.Println("..........................", usersResp)
	if err != nil {
		return &relation.FollowListResp{
			StatusCode: usersResp.StatusCode,
			StatusMsg:  usersResp.StatusMsg,
			UserList:   nil,
		}, err
	}

	// 数据库查询是否关注
	for _, item := range usersResp.Users {
		res, err := dal.IsFollow(item.Id, userId)
		if err == nil {
			item.IsFollow = res
		}
	}

	return &relation.FollowListResp{
		StatusCode: usersResp.StatusCode,
		StatusMsg:  usersResp.StatusMsg,
		UserList:   rpc.PackUsers(usersResp.Users),
	}, nil
}

// FollowerListMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerListMethod(ctx context.Context, request *relation.FollowerListReq) (resp *relation.FollowerListResp, err error) {
	// TODO: Your code here...
	followerIds, err := dal.GetFollowerList(request.UserId)
	if err != nil {
		return &relation.FollowerListResp{
			StatusCode: 1,
			StatusMsg:  "get follower ids fail",
			UserList:   nil,
		}, err
	}
	userResp, err := rpc.UserClient.InfosMethod(ctx, &user.InfosReq{UserIds: followerIds})
	if err != nil {
		return &relation.FollowerListResp{
			StatusCode: userResp.StatusCode,
			StatusMsg:  userResp.StatusMsg,
			UserList:   nil,
		}, err
	}

	// 数据库查询是否关注
	for _, item := range userResp.Users {
		res, err := dal.IsFollow(request.UserId, item.Id)
		if err == nil {
			item.IsFollow = res
		}
	}
	return &relation.FollowerListResp{
		StatusCode: userResp.StatusCode,
		StatusMsg:  userResp.StatusMsg,
		UserList:   rpc.PackUsers(userResp.Users),
	}, nil
}

// FriendListMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendListMethod(ctx context.Context, request *relation.FriendListReq) (resp *relation.FriendListResp, err error) {
	userId := request.UserId
	friendIds, err := dal.GetFriendList(userId)
	if err != nil {
		return &relation.FriendListResp{
			StatusCode: 1,
			StatusMsg:  "get friend id list fail",
			UserList:   nil,
		}, nil
	}
	fmt.Println(",,,,,,,,,,,,,,,,,,,,,,,", friendIds)
	usersResp, err := rpc.UserClient.InfosMethod(ctx, &user.InfosReq{UserIds: friendIds})
	if err != nil {
		return &relation.FriendListResp{
			StatusCode: 1,
			StatusMsg:  "get friend info list fail",
			UserList:   nil,
		}, nil
	}

	for _, item := range usersResp.Users {
		item.IsFollow = true
	}

	// TODO 获取聊天记录

	friends := rpc.PackFriendUsers(usersResp.Users)

	return &relation.FriendListResp{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   friends,
	}, nil
}

// IsFollowingMethod implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollowingMethod(ctx context.Context, request *relation.IsFollowingReq) (resp *relation.IsFollowingResp, err error) {
	userId := request.UserId
	toUserId := request.ToUserId

	res, err := dal.IsFollow(toUserId, userId)
	if err != nil {
		return &relation.IsFollowingResp{
			StatusCode:    1,
			StatusMsg:     "get relation ship fail",
			FollowingType: 0,
		}, err
	}
	var ftype int8
	if res {
		ftype = 1
	} else {
		ftype = 2
	}
	return &relation.IsFollowingResp{
		StatusCode:    0,
		StatusMsg:     "get relation ship success",
		FollowingType: ftype,
	}, nil
}
