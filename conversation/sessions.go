package conversation

import (
	"blog_0/configure"
	"blog_0/orm/user"
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

func SetSessionUser(context *gin.Context, us *user.User) {
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

func GetSessionUser(context *gin.Context) *user.User {
	UidString, err := context.Cookie(configure.SessionName)
	UidInt, err2 := strconv.Atoi(UidString)
	if err == nil && err2 == nil {
		u := user.User{
			Uid: UidInt,
		}
		u.GetUser()
		return &u
	} else {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.Cookie,
		})
	}
	return nil
}
