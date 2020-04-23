package conversation

import (
	"blog_0/configure"
	"blog_0/orm/user"
	"blog_0/proerror"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetSessionUser(context *gin.Context, us user.User) {
	Uid := strconv.Itoa(us.Uid)
	ck := &http.Cookie{
		Name:   configure.SessionName,
		Value:  Uid,
		Path:   "/",
		Domain: "*",
		MaxAge: 120 * 10,
	}
	http.SetCookie(context.Writer, ck)
}
func GetSessionUser(context *gin.Context) *user.User {
	UidString, err := context.Cookie(configure.SessionName)
	UidInt, err2 := strconv.Atoi(UidString)
	if err == nil && err2 == nil {
		u := user.User{
			Uid:        UidInt,
			User:       "",
			PassWord:   "",
			Type:       0,
			AvatarUrl:  "",
			CreateTime: nil,
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