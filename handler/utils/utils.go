package utils

import (
	"blog_0/configure"
	"blog_0/proerror"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bennyscetbun/jsongo"
	"github.com/gin-gonic/gin"
	"strings"
)

//提取出的公共接口，将对象转为json返回到前端，如果错误直接抛出
func SetSuccessRetObjectToJSONWithThrowException(context *gin.Context, obj interface{}) {

	bytes, err := json.Marshal(&obj)
	fmt.Println(string(bytes))
	if err == nil {
		context.Set(configure.ContextFiledName, string(bytes))
		return
	}
	panic(proerror.PanicError{
		ErrorType: proerror.ErrorOpera,
		ErrorCode: proerror.UnknownError,
	})
}

func SetFailedRetObjectToJSONWithThrowException(context *gin.Context, obj interface{}) {
	bytes, err := json.Marshal(&obj)
	if err == nil {
		context.Set(configure.ContextErrorFiledName, string(bytes))
		return
	}
	panic(proerror.PanicError{
		ErrorType: proerror.ErrorOpera,
		ErrorCode: proerror.UnknownError,
	})
}

func JsonParseWithThrowException(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.UnknownError,
		})
	}
	return string(bytes)
}

func JsonGoParseWithThrowException(node *jsongo.Node) string {
	r, err := json.MarshalIndent(node, "", "")
	if err != nil {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.UnknownError,
		})
	}
	return string(r)
}

func JsonGoUnmarshalToObjectWithThrowException(jsont string) jsongo.Node {
	r2 := jsongo.Node{}
	err := json.Unmarshal([]byte(jsont), &r2)
	if err != nil {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.UnknownError,
		})
	}
	return r2
}

//去除字符串中的回车符
func RemoveEnterChar(content string) string {
	return strings.Replace(content, "\n", "", -1)
}

func Base64String(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Decode64String(text string) (string, bool) {
	bs, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", false
	}
	return string(bs), true
}

func Decode64StringWithThrowException(text string) string {
	r, b := Decode64String(text)
	if b == false {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.UnknownError,
		})
	}
	return r
}
