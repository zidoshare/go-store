package common

import (
	"testing"
)

func TestToken(t *testing.T) {
	Conf = new(Configuration)
	Conf.Spwd = "123"
	Conf.Iss = "zido.site"
	Conf.LoginExp = 20000
	Conf.LogLevel = "info"
	token, err := NewToken(1, "user")
	if err != nil {
		t.Error(err)
	}
	tokenStr := token.String()
	parseToken := TokenFrom(tokenStr)
	if parseToken.Valid == false {
		t.Error("token is invalid")
	}
	if parseToken.Expired == true {
		t.Error("token is expired")
	}
	if parseToken.Payload.UID != token.Payload.UID {
		t.Error("token has't any data")
	}
}
