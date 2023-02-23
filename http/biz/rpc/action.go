package rpc

import (
	"douyinProject/http/biz/consts"
	"douyinProject/http/biz/model/http"
	"douyinProject/http/kitex_gen/httprpc"
	"douyinProject/http/kitex_gen/httprpc/actionservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var ActionClient actionservice.Client

func InitActionClient() {

	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}
	ActionClient, err = actionservice.NewClient(consts.ActionServiceName, client.WithResolver(r))
}

func PackComment(comm *httprpc.Comment) (*http.Comment, error) {

	user, err := PackUser(comm.User)
	if err != nil {
		return nil, fmt.Errorf("err")
	}
	var res = &http.Comment{
		ID:         comm.Id,
		User:       user,
		Content:    comm.Content,
		CreateDate: comm.CreateDate,
	}

	return res, nil
}

func PackComments(comms []*httprpc.Comment) ([]*http.Comment, error) {
	var cms = make([]*http.Comment, 0)
	for _, c := range comms {
		cmm, err := PackComment(c)
		if err != nil {
			return nil, err
		}
		cms = append(cms, cmm)
	}
	return cms, nil
}
