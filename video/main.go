package main

import (
	"douyinProject/video/cmd/consts"
	"douyinProject/video/cmd/dal"
	"douyinProject/video/cmd/minio"
	"douyinProject/video/cmd/rpc"
	video "douyinProject/video/kitex_gen/video/videoservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	rpc.InitUserService()
	rpc.InitRelationClient()
	rpc.InitActionCLient()
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddr}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	dal.InitDB()

	minio.InitMinioClient()

	addr, _ := net.ResolveTCPAddr("tcp", consts.VideoServiceHost+consts.VideoServicePort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
		server.WithRegistry(r),
	)

	svr := video.NewServer(new(VideoServiceImpl),
		opts...,
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
