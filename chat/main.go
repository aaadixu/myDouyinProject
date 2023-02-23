package main

import (
	"douyinProject/chat/cmd/consts"
	"douyinProject/chat/cmd/dal"
	"douyinProject/chat/cmd/redis"
	chat "douyinProject/chat/kitex_gen/chat/chatservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	dal.InitDB()
	redis.InitRedis()
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddr}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", consts.ChatServiceHost+consts.ChatServicePort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr),

		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ChatServiceName}),
		server.WithRegistry(r),
	)

	svr := chat.NewServer(new(ChatServiceImpl), opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
