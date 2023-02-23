package rpc

import (
	"douyinProject/video/cmd/consts"
	"douyinProject/video/kitex_gen/action/actionservice"
	"douyinProject/video/kitex_gen/relation/relationservice"
	"douyinProject/video/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var UserClient userservice.Client

var RelationClient relationservice.Client

var ActionClient actionservice.Client

func InitUserService() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}
	UserClient, err = userservice.NewClient(consts.UserServiceName, client.WithResolver(r))
}

func InitRelationClient() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}

	RelationClient, err = relationservice.NewClient(consts.RelationServiceName, client.WithResolver(r))
}

func InitActionCLient() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}
	ActionClient, err = actionservice.NewClient(consts.ActionServiceName, client.WithResolver(r))
}
