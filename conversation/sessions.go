package conversation

import (
	"blog_0/configure"
	"blog_0/orm/userDao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//删除用户会话  用在注销的时候
func SessionDestroy(context *gin.Context) {
	us := GetSessionUser(context)
	if us != nil {
		Uid := strconv.Itoa(us.Uid)
		ck := &http.Cookie{
			Name:   configure.SessionName,
			Value:  Uid,
			Path:   "/",
			Domain: "*",
			MaxAge: 1,
		}
		http.SetCookie(context.Writer, ck)
	}
}

//设置用户会话 在登录时调用
func SetSessionUser(context *gin.Context, us *userDao.User) {
	Uid := strconv.Itoa(us.Uid)
	ck := &http.Cookie{
		Name:   configure.SessionName,
		Value:  Uid,
		Path:   "/",
		Domain: "*",
		MaxAge: 120 * 10 * 5,
	}
	http.SetCookie(context.Writer, ck)
}

//获取用户会话,用于请求前校验
func GetSessionUser(context *gin.Context) *userDao.User {
	UidString, err := context.Cookie(configure.SessionName)
	UidInt, err2 := strconv.Atoi(UidString)
	if err == nil && err2 == nil {
		u := userDao.User{
			Uid: UidInt,
		}
		u.QueryGetUser()
		return &u
	}
	return nil
}
