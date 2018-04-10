package common

import (
	"errors"
	"fmt"
)

//ErrFooTokenInvalid token is invalid
var ErrFooTokenInvalid = errors.New("token is invalid")

//ErrFooBadTokenData token contains bad data
var ErrFooBadTokenData = errors.New("token contains bad data")

//ErrFooTokenExpired token expired
var ErrFooTokenExpired = errors.New("token expired")

//RespErr the REST response error model
type RespErr struct {
	Code       Code   `json:"code"`       //error code
	Message    string `json:"message"`    //base message
	Additional string `json:"additional"` //Additional Information
}

func (err *RespErr) String() string {
	if err.Additional != "" {
		return fmt.Sprintf("%s[Additional:%s][Code:%d]", err.Message, err.Additional, err.Code)
	}
	return fmt.Sprintf("%s[Code:%d]", err.Message, err.Code)
}
