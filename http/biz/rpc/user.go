package rpc

import (
	"douyinProject/http/biz/consts"
	"douyinProject/http/biz/model/http"
	"douyinProject/http/kitex_gen/httprpc"
	"douyinProject/http/kitex_gen/httprpc/userservice"
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

func PackUser(u *httprpc.User) (*http.User, error) {
	var us = &http.User{
		ID:              u.Id,
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
	return us, nil
}

func PackUsers(us []*httprpc.User) ([]*http.User, error) {
	var res = make([]*http.User, 0)
	for _, item := range us {
		uu, _ := PackUser(item)
		res = append(res, uu)
	}
	return res, nil
}
