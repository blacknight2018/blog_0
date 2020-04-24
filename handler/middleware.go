package handler

import (
	"blog_0/configure"
	"blog_0/logger"
	"blog_0/proerror"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func setCors(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", configure.AllowHttpServerCorAddress)
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	context.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
	context.Header("Access-Control-Allow-Credentials", "true")
}

/* 检查登录状态 */
func CheckLoginStatus(context *gin.Context) {
	//检查是否已经登录
	//conversation.GetSessionUser(context)
	context.Next()
}

/* 处理请求前后 */
func RequestMiddle(context *gin.Context) {

	//允许跨域
	setCors(context)
	//
	context.Next()
	//

	defer func() {
		if err := recover(); err != nil {
			logger.SimpleLog()
		}
	}()

	var obj = make(map[string]interface{})

	//检查是否有错误发生
	resp, err := context.Get(configure.ContextFiledName)
	if err {
		//对响应的JSON添加更多的字段
		obj[configure.ContextFiledName] = resp
		obj[configure.ResponseStatusFiledName] = "ok"
	} else {
		resp2, err2 := context.Get(configure.ContextErrorFiledName)
		if err2 {
			obj[configure.ResponseStatusFiledName] = "error"
			obj[configure.ContextErrorFiledName] = resp2
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
		}
	}

	//
	output, err3 := json.Marshal(&obj)
	if err3 == nil {
		context.Writer.WriteString(string(output))
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

/* 处理异常 */
func Except(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			error, ok := err.(proerror.PanicError)
			if ok {
				switch error.ErrorType {
				case proerror.ErrorIo:
					//记录日志
					logger.SimpleLog()
				case proerror.ErrorOpera:
					//业务错误
					//json, _ := json.Marshal(&err)
					//context.Writer.WriteString(string(json))
				}
				context.Set(configure.ContextErrorFiledName, err)
				//阻止传播给下层
				context.Abort()
			} else {
				//记录日志
				logger.SimpleLog()
			}
		}
	}()
	context.Next()
}
