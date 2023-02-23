package dal

import (
	"context"
	"douyinProject/action/cmd/rpc"
	"douyinProject/action/kitex_gen/action"
	"douyinProject/action/kitex_gen/user"
)

// 新增一条评论
func AddComment(userId, videoId int64, content *string) (*action.Comment, error) {
	var comm = &Comment{
		VideoId: videoId,
		UserId:  userId,
		Content: *content,
	}
	res := DB.Create(&comm)
	if res.Error != nil {
		return nil, res.Error
	}

	t := comm.CreatedAt

	time := t.Format("01-02")

	return &action.Comment{
		Id:         int64(comm.ID),
		User:       nil,
		Content:    *content,
		CreateDate: time,
	}, nil

}

// 根据评论id删除评论（逻辑删除）
func DeleteComment(vid int64, comId int64) error {
	res := DB.Where("id = ? and video_id = ?", comId, vid).Delete(&Comment{})
	return res.Error
}

// 根据视频id获取所有评论
func GetCommentList(videoId int64) ([]*action.Comment, error) {
	var comms []Comment
	DB.Where("video_id = ?", videoId).Find(&comms)
	var res = make([]*action.Comment, 0)
	for _, c := range comms {
		authorId := c.UserId
		packComm, _ := PackComment(c)
		// 获取用户信息(视频作者)
		resp, err := rpc.UserClient.InfoMethod(context.Background(), &user.InfoReq{UserId: authorId})
		if err != nil {
			return nil, err
		}
		user, _ := PackUser(resp.User)
		packComm.User = user
		res = append(res, packComm)
	}

	return res, nil
}

// 根据评论id查询评论的作者id
func GetCommentAuthorId(commentId int64) (int64, error) {
	var comm Comment
	res := DB.Where("id = ?", commentId).Find(&comm)
	if res.Error != nil {
		return -1, res.Error
	}
	return comm.UserId, nil

}
