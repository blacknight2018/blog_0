package handler

import (
	"blog_0/configure"
	"blog_0/conversation"
	"blog_0/orm"
	"blog_0/proerror"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func init() {
	gob.Register(&orm.User{})
}
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
func UserLogin(context *gin.Context) {
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
		conversation.SetSessionUser(context, us)
		context.Set(configure.ContextFiledName, us)

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func UserQuery(context *gin.Context) {
	us := conversation.GetSessionUser(context)
	if us != nil {
		//fmt.Println(us)
		us.Get()
		context.Set(configure.ContextFiledName, us)
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
