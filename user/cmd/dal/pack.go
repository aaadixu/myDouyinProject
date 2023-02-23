package dal

import "douyinProject/user/kitex_gen/user"

func PackUser(userDB *User) *user.User {
	if userDB == nil {
		return &user.User{}
	}
	return &user.User{
		Id:              int64(userDB.ID),
		Name:            userDB.Username,
		FollowCount:     userDB.FollowCount,
		FollowerCount:   userDB.FollowerCount,
		IsFollow:        userDB.IsFollow,
		Avatar:          userDB.Avatar,
		BackgroundImage: userDB.BackgroundImage,
		Signature:       userDB.Signature,
		TotalFavorited:  userDB.TotalFavorited,
		WorkCount:       userDB.WorkCount,
		FavoriteCount:   userDB.FavoriteCount,
	}
}

func PackUsers(usersDB []*User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range usersDB {
		user2 := PackUser(u)
		users = append(users, user2)
	}
	return users
}
