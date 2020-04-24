package proerror

/* 业务错误码 */
const (
	ParamError   = iota
	UnknownError = iota
	LoginError   = iota
	LoginBefore  = iota
)

/* 错误类型 */
const (
	ErrorOpera = iota
)

type PanicError struct {
	ErrorType int
	ErrorCode int
}
