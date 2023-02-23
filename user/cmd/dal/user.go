package dal

import "douyinProject/user/kitex_gen/user"

func ExistUserName(username string) bool { // 注册检测
	users := make([]User, 0)
	DB.Where("user_name = ?", username).Find(&users)
	return len(users) == 0 // 如果返回true，说明用户名不存在，可以注册
}

func CheckUser(username, password string) (*user.User, error) { // 登录检测
	var user User
	res := DB.Where("user_name = ? and password = ?", username, password).Find(&user)
	return PackUser(&user), res.Error
}

func CreateUser(user User) (*user.User, error) { // 注册用户
	res := DB.Create(&user)
	return PackUser(&user), res.Error
}

func UserInfo(id int64) (*user.User, error) { // 用户信息
	var us User
	res := DB.Where("id = ?", id).Find(&us)
	return PackUser(&us), res.Error
}

func UserInfos(ids []int64) ([]*user.User, error) {
	users := make([]*User, 0)
	res := DB.Where("id IN ?", ids).Find(&users)
	return PackUsers(users), res.Error
}

func UserAddWork(id int64) error {
	var user User
	res := DB.Where("id = ?", id).Find(&user)
	if res.Error != nil {
		return res.Error
	}
	user.WorkCount = user.WorkCount + 1
	res = DB.Save(&user)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func FavoriteCount(uid int64, count int8) error {
	var user User
	res := DB.Where("id = ?", uid).Find(&user)
	if res.Error != nil {
		return res.Error
	}
	user.FavoriteCount = user.FavoriteCount + int64(count)
	res = DB.Save(&user)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func TotalFavorited(uid int64, count int8) error {
	var user User
	res := DB.Where("id = ?", uid).Find(&user)
	if res.Error != nil {
		return res.Error
	}
	user.TotalFavorited = user.TotalFavorited + int64(count)

	res = DB.Save(&user)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func FollowCount(uid int64, count int64) error {
	var user User
	res := DB.Where("id = ?", uid).Find(&user)
	if res.Error != nil {
		return res.Error
	}
	user.FollowCount = user.FollowCount + count
	res = DB.Save(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func FollowerCount(uid int64, count int64) error {
	var user User
	res := DB.Where("id = ?", uid).Find(&user)
	if res.Error != nil {
		return res.Error
	}
	user.FollowerCount = user.FollowerCount + count
	res = DB.Save(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
