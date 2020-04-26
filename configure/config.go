package configure

//Handler Const String
const ContextFiledName = "data"
const ContextErrorFiledName = "error"
const ContextEmptyFiled = ""
const ResponseSuccessName = "ok"
const ResponseErrorName = "error"
const ResponseStatusFiledName = "status"
const AllowHttpServerCorAddress = "http://127.0.0.1"
const SelectFiledKeyName = "filed"
const EmptyString = ""
const SessionName = "Blog"
const (
	U_ROOT   = iota
	U_ADMIN_ = iota
	U_USER   = iota
)

//数据库链接
func GetDSN() string {
	return `root:root@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`
}

//文件保存的路径
func GetLocalFileDir() string {
	return `./files/`
}
