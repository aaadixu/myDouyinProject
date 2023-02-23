package rpc

import (
	"douyinProject/action/cmd/consts"
	"douyinProject/action/kitex_gen/user/userservice"
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
