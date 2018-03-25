package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

//TokenHeader token header
type TokenHeader struct {
	Typ string `json:"typ"`   //token type: JWT(JSON Web Token)
	Alg string `json:"HS256"` //secure method
}

//TokenPayload token payload
type TokenPayload struct {
	Iss  string    `json:"iss"`  //issuer
	Exp  time.Time `json:"exp"`  //expiration time
	UID  uint      `json:"uid"`  //user id
	Role string    `json:"role"` // user role name
}

const (
	tokenSep = "."
)

var (
	//ErrFooTokenInvalid token is invalid error
	ErrFooTokenInvalid = errors.New("token is invalid")
	//ErrFooBadTokenData token is bad data
	ErrFooBadTokenData = errors.New("token is bad data")
)

var Cipher = hmac.New(sha256.New, []byte(Conf.Spwd))

//NewToken create default token
func NewToken(uid uint, role string) (*Token, error) {
	now := time.Now()
	header := &TokenHeader{
		Typ: "JWT",
		Alg: "HS256",
	}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		logger.Errorf("create token : build header : json marshal token header failed:%+v", header)
		return nil, ErrFooBadTokenData
	}
	exp := now.Add(time.Duration(Conf.LoginExp) * time.Second)
	payload := &TokenPayload{
		Iss:  Conf.Iss,
		Exp:  exp,
		UID:  uid,
		Role: role,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		logger.Errorf("create token : build payload : json marshal token payload failed:%+v", payload)
		return nil, ErrFooBadTokenData
	}

	headerStr := base64.StdEncoding.EncodeToString(headerJSON)
	payloadStr := base64.StdEncoding.EncodeToString(payloadJSON)

	preSign := headerStr + tokenSep + payloadStr

	tailStr := base64.StdEncoding.EncodeToString(Cipher.Sum([]byte(preSign)))
	Cipher.Reset()
	token := &Token{
		header:     header,
		Payload:    payload,
		HeaderStr:  headerStr,
		PayloadStr: payloadStr,
		Sign:       tailStr,
		Valid:      true,
		Expired:    false,
	}
	return token, nil
}

//Token token obj
type Token struct {
	header     *TokenHeader // Will not resolve the first time,need call func getHeader
	HeaderStr  string
	Payload    *TokenPayload // Will resolve the first time
	PayloadStr string
	Sign       string
	Valid      bool //is/not valid (whether expired or not)
	Expired    bool // is expired
}

func (token *Token) String() string {
	return token.HeaderStr + "." + token.PayloadStr + "." + token.Sign
}

func (token *Token) getHeader() *TokenHeader {
	if token.header == nil {
		token.header = &TokenHeader{
			Typ: "JWT",
			Alg: "HS256",
		}
	}
	return token.header
}

//UpdateSign update expired by now
func (token *Token) UpdateSign() {
	token.Payload.Exp = time.Now().Add(time.Duration(Conf.LoginExp) * time.Second)
	token.BuildSign()
}

//BuildSign build sign
func (token *Token) BuildSign() error {
	headerJSON, err := json.Marshal(token.getHeader())
	if err != nil {
		logger.Errorf("create token : build header : json marshal token header failed:%+v", token.header)
		return ErrFooBadTokenData
	}
	payloadJSON, err := json.Marshal(token.Payload)
	if err != nil {
		logger.Errorf("create token : build payload : json marshal token payload failed:%+v", token.Payload)
		return ErrFooBadTokenData
	}

	headerStr := base64.StdEncoding.EncodeToString(headerJSON)
	payloadStr := base64.StdEncoding.EncodeToString(payloadJSON)
	preSign := headerStr + tokenSep + payloadStr
	token.Sign = base64.StdEncoding.EncodeToString(Cipher.Sum([]byte(preSign)))
	return nil
}

//TokenFrom parse token from str
func TokenFrom(tokenStr string) (token *Token) {
	token = &Token{}
	if tokenStr == "" {
		token.Valid = false
		return
	}
	arr := strings.Split(tokenStr, tokenSep)
	if len(arr) != 3 {
		token.Valid = false
		return
	}
	token.HeaderStr = arr[0]
	token.PayloadStr = arr[1]
	token.Sign = arr[2]
	if token.HeaderStr == "" || token.PayloadStr == "" || token.Sign == "" {
		token.Valid = false
		return
	}
	rebuildSign := Cipher.Sum([]byte(token.HeaderStr + tokenSep + token.PayloadStr))
	if base64.StdEncoding.EncodeToString(rebuildSign) != token.Sign {
		token.Valid = false
		return
	}

	token.Valid = true
	var err error
	token.Payload, err = ParsePayload(token.PayloadStr)
	if err != nil {
		return
	}
	exp := token.Payload.Exp
	if exp.After(time.Now()) {
		token.Expired = false
	} else {
		token.Expired = true
	}
	return
}

//ParsePayload parse payload from PayloadStr without decoding by base64
func ParsePayload(PayloadStr string) (payload *TokenPayload, err error) {
	payload = &TokenPayload{}
	jsonStr, dErr := base64.StdEncoding.DecodeString(PayloadStr)
	if err != nil {
		err = dErr
		return
	}
	err = json.Unmarshal(jsonStr, payload)
	return
}