package rpc

import (
	"douyinProject/relation/cmd/consts"
	"douyinProject/relation/kitex_gen/relation"
	"douyinProject/relation/kitex_gen/user"
	"douyinProject/relation/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var UserClient userservice.Client

func InitUserClient() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}
	UserClient, err = userservice.NewClient(consts.UserServiceName, client.WithResolver(r))
}

func PackUser(u *user.User) *relation.User {
	return &relation.User{
		Id:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowerCount,
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

func PackUsers(us []*user.User) []*relation.User {
	var res = make([]*relation.User, 0)
	for _, item := range us {
		res = append(res, PackUser(item))
	}
	return res
}

func PackFriendUser(u *user.User) *relation.FriendUser {
	return &relation.FriendUser{
		Id:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowerCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
		Message:         "",
		MsgType:         0,
	}
}

func PackFriendUsers(us []*user.User) []*relation.FriendUser {
	var res = make([]*relation.FriendUser, 0)
	for _, item := range us {
		res = append(res, PackFriendUser(item))
	}
	return res
}
