package rpc

import (
	"douyinProject/action/cmd/consts"
	"douyinProject/action/kitex_gen/video/videoservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var VideoClient videoservice.Client

func InitVideoClient() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}
	VideoClient, err = videoservice.NewClient(consts.VideoServiceName, client.WithResolver(r))
}
