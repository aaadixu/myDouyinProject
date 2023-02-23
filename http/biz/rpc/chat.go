package rpc

import (
	"douyinProject/http/biz/consts"
	"douyinProject/http/biz/model/http"
	"douyinProject/http/kitex_gen/httprpc"
	"douyinProject/http/kitex_gen/httprpc/chatservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var ChatClient chatservice.Client

func InitChatClient() {

	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddr})
	if err != nil {
		log.Fatal(err)
	}
	ChatClient, err = chatservice.NewClient(consts.ChatServiceName, client.WithResolver(r))
	if err != nil {
		fmt.Errorf("connect chat server err")
	}
}

func PackMessage(m *httprpc.Message) *http.Message {
	return &http.Message{
		ID:         m.Id,
		ToUserID:   m.ToUserId,
		FromUserID: m.ToUserId,
		Content:    m.Content,
		CreateTime: &m.CreateTime,
	}
}

func Packmessages(ms []*httprpc.Message) []*http.Message {
	var res = make([]*http.Message, 0)
	for _, item := range ms {
		res = append(res, PackMessage(item))
	}
	return res
}
