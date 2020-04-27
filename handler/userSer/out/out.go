package out

import "blog_0/orm/userDao"

func GetUserName(Uid int) (string, bool) {
	u := userDao.User{
		Uid: Uid,
	}
	if false == u.QueryGetUser() {
		return "", false
	}
	return u.User, true
}
func GetUserAvatar(Uid int) (string, bool) {
	u := userDao.User{
		Uid: Uid,
	}
	if false == u.QueryGetUser() {
		return "", false
	}
	return u.AvatarUrl, true
}
