package dal

import (
	"douyinProject/chat/kitex_gen/chat"
)

func PackMessage(msg ChatRecord) *chat.Message {
	return &chat.Message{
		Id:         int64(msg.ID),
		ToUserId:   msg.ToUserId,
		FromUserId: msg.FromUserId,
		Content:    msg.Content,
		CreateTime: msg.CreateTime,
	}
}

func PackMessages(msgs []*ChatRecord) []*chat.Message {
	var res = make([]*chat.Message, 0)
	for _, msg := range msgs {
		res = append(res, PackMessage(*msg))
	}
	return res
}
