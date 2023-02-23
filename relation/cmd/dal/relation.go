package dal

func AddFollow(followId, followerId int64) error {

	var follow = &Follow{
		FollowId:   followId,
		FollowerId: followerId,
	}
	res := DB.Create(&follow)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteFollow(followId, followerId int64) error {
	var follow = &Follow{
		FollowId:   followId,
		FollowerId: followerId,
	}
	res := DB.Where("follow_id = ? and follower_id = ?", followId, followerId).Delete(&follow)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func IsFollow(followId, followerId int64) (bool, error) {
	var follow Follow

	res := DB.Where("follow_id = ? ", followId).Where("follower_id = ?", followerId).Find(&follow)

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil

}

func AddFriend(u1, u2 int64) error {

	var friend = &Friend{
		User1Id: u1,
		User2Id: u2,
	}
	res := DB.Create(&friend)
	return res.Error
}

func DeleteFriend(u1, u2 int64) error {
	var friend1 Friend

	res := DB.Where("user1_id = ?", u1).Where("user2_id = ?", u2).Delete(&friend1)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return nil
	}
	res = DB.Where("user1_id = ?", u2).Where("user2_id = ?", u1).Delete(&friend1)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return nil
	}
	return nil
}

func ExistedFriendship(u1, u2 int64) bool {
	var f Friend
	res := DB.Where("user1_id = ?", u1).Where("user2_id = ?", u2).Find(&f)
	if res.Error != nil {
		return false
	}
	if res.RowsAffected > 0 {
		return true
	}
	res = DB.Where("user1_id = ?", u2).Where("user2_id = ?", u1).Find(&f)
	if res.Error != nil {
		return false
	}
	if res.RowsAffected > 0 {
		return true
	}
	return false
}

func GetFollowList(uid int64) ([]int64, error) {
	var follows []Follow
	var ids = make([]int64, 0)
	res := DB.Where("follower_id = ?", uid).Find(&follows)
	if res.Error != nil {
		return ids, res.Error
	}
	for _, item := range follows {
		ids = append(ids, item.FollowId)
	}
	return ids, nil
}

func GetFollowerList(uid int64) ([]int64, error) {
	var followers []Follow
	var ids = make([]int64, 0)
	res := DB.Where("follow_id = ?", uid).Find(&followers)
	if res.Error != nil {
		return ids, res.Error
	}
	for _, item := range followers {
		ids = append(ids, item.FollowerId)
	}
	return ids, nil
}

func GetFriendList(uid int64) ([]int64, error) {
	var friend1s []Friend
	var friend2s []Friend
	var ids = make([]int64, 0)

	res := DB.Where("user1_id = ?", uid).Find(&friend1s)
	if res.Error != nil {
		return ids, res.Error
	}
	for _, item := range friend1s {
		ids = append(ids, item.User2Id)
	}
	res = DB.Where("user2_id = ?", uid).Find(&friend2s)
	if res.Error != nil {
		return ids, res.Error
	}
	for _, item := range friend2s {
		ids = append(ids, item.User1Id)
	}
	return ids, nil
}
