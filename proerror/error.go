package proerror

/* 业务错误码 */
const (
	FieldEmpty = iota
	ParamError = iota
	FileEmpty  = iota
)

/* 错误类型 */
const (
	ErrorIo    = iota
	ErrorOpera = iota
)

type PanicError struct {
	ErrorType int
	ErrorCode int
}
