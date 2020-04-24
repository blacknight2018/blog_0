package conversation

import (
	"blog_0/configure"
	"blog_0/orm/userDao"
	"blog_0/proerror"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var onlineUser = make(map[int]bool)

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

func GetSessionUser(context *gin.Context) *userDao.User {
	UidString, err := context.Cookie(configure.SessionName)
	UidInt, err2 := strconv.Atoi(UidString)
	if err == nil && err2 == nil {
		u := userDao.User{
			Uid: UidInt,
		}
		u.QueryGetUser()
		return &u
	} else {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.NotLogin,
		})
	}
	return nil
}
