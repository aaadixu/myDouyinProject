package main

import (
	"douyinProject/relation/cmd/consts"
	"douyinProject/relation/cmd/dal"
	"douyinProject/relation/cmd/rpc"
	relation "douyinProject/relation/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {

	rpc.InitUserClient()

	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddr}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	dal.InitDB()

	addr, _ := net.ResolveTCPAddr("tcp", consts.RelationServiceHost+consts.RelationServicePort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.RelationServiceName}),
		server.WithRegistry(r),
	)

	svr := relation.NewServer(new(RelationServiceImpl), opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
