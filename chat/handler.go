package main

import (
	"context"
	"douyinProject/chat/cmd/dal"
	"douyinProject/chat/cmd/redis"
	chat "douyinProject/chat/kitex_gen/chat"
	"encoding/json"
	"strconv"
	"time"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// ChatMethod implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ChatMethod(ctx context.Context, request *chat.ChatReq) (resp *chat.ChatResp, err error) {
	userId := request.UserId
	toUserId := request.ToUserId
	preMsgTime := request.PreMsgTime
	// 当前接收其他用户，因此是别人to我
	key := strconv.FormatInt(toUserId, 10) + "to" + strconv.FormatInt(userId, 10)
	var resMsgs = make([]*chat.Message, 0)
	for {
		var msg []byte
		err = redis.RedisCli.LPop(key).Scan(&msg) // 从redis中弹出一条数据

		var resMsg dal.ChatRecord
		err = json.Unmarshal(msg, &resMsg) // 反序列化

		curMsgTime, _ := strconv.ParseInt(resMsg.CreateTime, 10, 64) // 判断时间
		if err != nil {
			break
		}
		if preMsgTime < curMsgTime {
			resMsgs = append(resMsgs, dal.PackMessage(resMsg)) // 满足时间条件则返回
		}
	}
	return &chat.ChatResp{
		StatusCode:  0,
		StatusMsg:   "success",
		MessageList: resMsgs,
	}, nil
}

// ActionMethod implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ActionMethod(ctx context.Context, request *chat.ActionReq) (resp *chat.ActionResp, err error) {
	fromUserId := request.UserId
	toUserId := request.ToUserId
	//actionType := request.ActionType // 未操作
	content := request.Content
	// 获取聊天记录
	// 先从redis里边取

	key := strconv.FormatInt(fromUserId, 10) + "to" + strconv.FormatInt(toUserId, 10)
	createTime := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 保存到数据库
	msg, err := dal.AddRecord(fromUserId, toUserId, content, createTime)
	if err != nil {
		return &chat.ActionResp{
			StatusCode: 1,
			StatusMsg:  "send message fail",
		}, err
	}
	data, err := json.Marshal(msg)
	_, err = redis.RedisCli.RPush(key, data).Result()
	if err != nil {
		return &chat.ActionResp{
			StatusCode: 2,
			StatusMsg:  "send message fail",
		}, err
	}

	return &chat.ActionResp{
		StatusCode: 0,
		StatusMsg:  "send message success",
	}, nil
}
