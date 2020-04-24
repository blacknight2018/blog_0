package handler

import (
	"blog_0/configure"
	"blog_0/conversation"
	"blog_0/orm/userDao"
	"blog_0/proerror"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"strconv"
)

func init() {
	gob.Register(&userDao.User{})
}
func InsertUser(context *gin.Context) {
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		user := gjson.Get(json, "user").String()
		password := gjson.Get(json, "password").String()
		avatar := gjson.Get(json, "avatar").String()

		us := userDao.User{
			User:      user,
			PassWord:  password,
			AvatarUrl: avatar,
		}
		us.InsertUser()
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func DeleteUserLogout(context *gin.Context) {
	conversation.SessionDestroy(context)
	context.Set(configure.ContextFiledName, configure.ContextEmptyFiled)
}

func InsertUserLogin(context *gin.Context) {
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		user := gjson.Get(json, "user").String()
		password := gjson.Get(json, "password").String()
		us := userDao.User{
			User:     user,
			PassWord: password,
		}
		us.QueryCheckUser()
		conversation.SetSessionUser(context, &us)
		context.Set(configure.ContextFiledName, us)

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func QueryUser(context *gin.Context) {
	//默认查询自己，
	us := conversation.GetSessionUser(context)
	uid := context.Query("uid")
	uidInt, err := strconv.Atoi(uid)
	if uid != "" && err == nil {
		us.Uid = uidInt
	}
	if us != nil {
		us.QueryGetUser()
		context.Set(configure.ContextFiledName, us)
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
