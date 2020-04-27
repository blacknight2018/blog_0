package configure

import (
	"blog_0/logger"
	"fmt"
	"os"
	"strings"
)

//Handler Const String
const ContextFiledName = "data"
const ContextErrorFiledName = "error"
const ContextEmptyFiled = ""
const ResponseSuccessName = "ok"
const ResponseErrorName = "error"
const ResponseStatusFiledName = "status"

var AllowHttpServerCorAddress string = "http://127.0.0.1"

const SelectFiledKeyName = "filed"
const EmptyString = ""
const SessionName = "Blog"
const (
	U_ROOT   = iota
	U_ADMIN_ = iota
	U_USER   = iota
)

func init() {
	fmt.Printf(os.Hostname())
	pcName, err := os.Hostname()
	if err == nil {
		if strings.Index(pcName, "Za631y") >= 0 {
			AllowHttpServerCorAddress = "http://39.107.93.78"
		} else {
			AllowHttpServerCorAddress = "http://127.0.0.1"
		}
	} else {
		logger.SimpleLog()
	}

}

//数据库链接
func GetDSN() string {
	return `root:root@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`
}

//文件保存的路径
func GetLocalFileDir() string {
	return `./files/`
}
