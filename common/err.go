package common

var (
	//ErrorFormatInvalid text invalid
	ErrorFormatInvalid = &RespError{
		Message: "格式错误",
	}
)

//RespError respond on error
type RespError struct {
	Message string `json:"message"`
}
