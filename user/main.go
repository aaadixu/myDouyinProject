package main

import (
	"douyinProject/user/cmd/consts"
	"douyinProject/user/cmd/dal"
	user "douyinProject/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddr}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	dal.InitDB()
	addr, _ := net.ResolveTCPAddr("tcp", consts.UserServiceHost+consts.UserServicePort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
		server.WithRegistry(r),
	)

	svr := user.NewServer(new(UserServiceImpl),
		opts...,
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
