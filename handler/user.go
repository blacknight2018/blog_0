package handler

import (
	"blog_0/configure"
	"blog_0/orm"
	"blog_0/proerror"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func UserInsert(context *gin.Context) {
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		user := gjson.Get(json, "user").String()
		password := gjson.Get(json, "password").String()
		avatar := gjson.Get(json, "avatar").String()
		us := orm.User{
			User:      user,
			PassWord:  password,
			AvatarUrl: avatar,
		}
		us.InsertUser()
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func setCookie(context *gin.Context, us orm.User) {
	session := sessions.Default(context)
	session.Set("user", us.User)
	session.Save()
}
func UserLogin(context *gin.Context) {
	var ret string
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		user := gjson.Get(json, "user").String()
		password := gjson.Get(json, "password").String()
		us := orm.User{
			User:     user,
			PassWord: password,
		}
		us.CheckUser()
		setCookie(context, us)
		context.Set(configure.ContextFiledName, ret)

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
