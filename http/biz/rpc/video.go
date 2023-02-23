package rpc

import (
	"douyinProject/http/biz/consts"
	"douyinProject/http/biz/model/http"
	"douyinProject/http/kitex_gen/httprpc"
	"douyinProject/http/kitex_gen/httprpc/videoservice"
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

func PackVideo(v *httprpc.Video) (*http.Video, error) {
	var rv http.Video
	rv.CommentCount = v.CommentCount

	rv.ID = int64(v.Id)
	rv.FavoriteCount = v.FavoriteCount
	rv.Title = v.Title
	rv.PlayURL = v.PlayUrl
	rv.CoverURL = v.CoverUrl
	rv.IsFavorite = v.IsFavorite

	author, _ := PackUser(v.Author)

	rv.Author = author
	return &rv, nil
}

func PackVideos(videos []*httprpc.Video) ([]*http.Video, error) {
	vs := make([]*http.Video, 0)
	for _, v := range videos {

		rpcv, _ := PackVideo(v)

		vs = append(vs, rpcv)
	}

	return vs, nil
}
