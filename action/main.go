package main

import (
	"douyinProject/action/cmd/consts"
	"douyinProject/action/cmd/dal"
	"douyinProject/action/cmd/rpc"
	action "douyinProject/action/kitex_gen/action/actionservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	dal.InitDB()
	rpc.InitUserClient()
	rpc.InitVideoClient()
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddr}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", consts.ActionServiceHost+consts.ActionServicePort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ActionServiceName}),
		server.WithRegistry(r),
	)

	svr := action.NewServer(new(ActionServiceImpl),
		opts...,
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
