package configure

const ContextFiledName = "data"

//数据库链接
func GetDSN() string {
	return `root:root@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`
}
