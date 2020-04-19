package handler

import (
	"blog_0/configure"
	"blog_0/logger"
	"blog_0/proerror"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func setCors(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	context.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
}

/* 处理请求前后 */
func RequestMiddle(context *gin.Context) {
	//允许跨域
	setCors(context)
	//

	//检查是否已经登录

	//

	context.Next()

	//运行到这里表面没有异常和错误
	resp, err := context.Get(configure.ContextFiledName)
	if err {
		var obj = make(map[string]interface{})
		//对响应的JSON添加更多的字段
		obj[configure.ContextFiledName] = resp
		//
		output, err := json.Marshal(&obj)
		if err == nil {
			context.Writer.WriteString(string(output))
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
		}
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
					json, _ := json.Marshal(&err)
					context.Writer.WriteString(string(json))
				}
				//阻止传播
				context.Abort()
			} else {
				//记录日志
				logger.SimpleLog()
			}
		}
	}()
	context.Next()
}
