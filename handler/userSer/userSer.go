package userSer

import (
	"blog_0/configure"
	"blog_0/handler/userSer/out/conversation"
	"blog_0/handler/utils"
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

type swag1 struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

// @添加用户
// Name will print hello name
// @Summary 添加用户
// @Description 添加用户
// @Accept json
// @Produce  json
// @Param u body swag1 true "提交的参数"
// @Router /user/add [POST]
// @Success 200
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
		if !us.InsertUser() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
	}
}

// @用户注销
// Name will print hello name
// @Summary 用户注销
// @Description 用户注销
// @Accept json
// @Produce  json
// @Router /user/logout [DELETE]
// @Success 200
func DeleteUserLogout(context *gin.Context) {
	conversation.SessionDestroy(context)
	utils.SetSuccessRetObjectToJSONWithThrowException(context, configure.ContextEmptyFiled)
}

type swag0 struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// @用户登录
// Name will print hello name
// @Summary 用户登录
// @Description 用户登录
// @Accept json
// @Produce  json
// @Param u body swag0 true "提交的参数"
// @Router /user/login [POST]
// @Success 200 {object} userDao.User
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
		if !us.QueryCheckUserPassWord() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.LoginError,
			})
		}
		conversation.SetSessionUser(context, &us)
		//context.Set(configure.ContextFiledName, us)
		utils.SetSuccessRetObjectToJSONWithThrowException(context, us)
	} else {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.UnknownError,
		})
	}
}

// @查询用户
// Name will print hello name
// @Summary 查询用户
// @Description 获取用户信息
// @Accept json
// @Produce  json
// @Param uid query int  false "用户uid"
// @Router /user [GET]
// @Success 200 {object} userDao.User
func QueryUser(context *gin.Context) {

	//默认查询自己，
	us := conversation.GetSessionUser(context)
	uid := context.Query("uid")
	uidInt, err := strconv.Atoi(uid)
	if us == nil {
		us = &userDao.User{}
	}
	if uid != "" && err == nil {
		us.Uid = uidInt
	}
	if !us.QueryGetUser() {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.UnknownError,
		})
	}
	utils.SetSuccessRetObjectToJSONWithThrowException(context, us)
}
