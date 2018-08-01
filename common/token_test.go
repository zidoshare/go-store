package common

import (
	"testing"
)

var _ = (func() interface{} {
	_testing = true
	return nil
}())

func TestParseToken(t *testing.T) {
	Conf = new(Configuration)
	Conf.Spwd = "123"
	Conf.Iss = "zido.site"
	Conf.LoginExp = 20000
	Conf.Alg = []string{"HS256"}
	token, err := NewToken(1, "user")
	if err != nil {
		t.Error(err)
	}
	tokenStr := token.String()
	parsedToken := TokenFrom(tokenStr)
	if parsedToken.Valid == false {
		t.Error("token is invalid")
	}
	if parsedToken.Expired == true {
		t.Error("token is expired")
	}
	if parsedToken.Payload.UID != token.Payload.UID {
		t.Error("token has't any data")
	}
	if parsedToken.Payload.UID != 1 {
		t.Error(("the data UID of token is missing"))
	}
	if parsedToken.Payload.Role != "user" {
		t.Error(("the data Role of token is missing"))
	}
}
