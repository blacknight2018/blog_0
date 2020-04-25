package handler

import (
	"blog_0/configure"
	"blog_0/conversation"
	"blog_0/handler/utils"
	"blog_0/logger"
	"blog_0/proerror"
	"github.com/bennyscetbun/jsongo"
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
	if nil == conversation.GetSessionUser(context) {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.LoginBefore,
		})
	}
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

	//ContextFiledName 域 必须是json
	resp, err := context.Get(configure.ContextFiledName)
	var ret jsongo.Node = jsongo.Node{}
	if err {
		respString := resp.(string)
		r2 := utils.JsonGoUnmarshalToObjectWithThrowException(respString)
		ret.At(configure.ContextFiledName).Val(r2)
		ret.At(configure.ResponseStatusFiledName).Val(configure.ResponseSuccessName)

	} else {
		resp2, err2 := context.Get(configure.ContextErrorFiledName)
		resp2String := resp2.(string)
		if err2 {
			r2 := utils.JsonGoUnmarshalToObjectWithThrowException(resp2String)
			ret.At(configure.ContextErrorFiledName).Val(r2)
			ret.At(configure.ResponseStatusFiledName).Val(configure.ResponseErrorName)
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
		}
	}
	context.Writer.WriteString(utils.JsonGoParseWithThrowException(&ret))
	context.Writer.Flush()
}

/* 处理异常 */
func Except(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			error, ok := err.(proerror.PanicError)
			if ok {
				switch error.ErrorType {

				case proerror.ErrorOpera:
					//业务错误
					//json, _ := json.Marshal(&err)
					//context.Writer.WriteString(string(json))
				}
				utils.SetFailedRetObjectToJSONWithThrowException(context, err)
				//阻止传播给下层
				context.Abort()
			}
			//记录日志
			logger.SimpleLog()
		}
	}()
	context.Next()
}
