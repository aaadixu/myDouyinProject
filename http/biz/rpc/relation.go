package rpc

import (
	"douyinProject/http/biz/consts"
	"douyinProject/http/biz/model/http"
	"douyinProject/http/kitex_gen/httprpc"
	"douyinProject/http/kitex_gen/httprpc/relationservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var RelationClient relationservice.Client

func InitRelationClient() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}

	RelationClient, err = relationservice.NewClient(consts.RelationServiceName, client.WithResolver(r))

}

func PackFrined(u *httprpc.FriendUser) *http.FriendUser {
	return &http.FriendUser{
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
		Message:         u.Message,
		MsgType:         u.MsgType,
	}
}

func PackFriends(us []*httprpc.FriendUser) []*http.FriendUser {
	res := make([]*http.FriendUser, 0)
	for _, item := range us {
		res = append(res, PackFrined(item))
	}
	return res
}
