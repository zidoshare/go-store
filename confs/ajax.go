package confs

//Code reponse code value
type Code int

const (
	//SUCCESS code
	SUCCESS Code = 0
)

//Resp the REST response model
type Resp struct {
	code Code
	data *interface{}
}

//NewResp new response model
func NewResp() *Resp {
	return &Resp{}
}

//Success new response model with Success
func Success(data *interface{}) *Resp {
	return &Resp{0, data}
}

//Fail new response model with Fail,if code is 0,reset it to -1
func Fail() *Resp {
	return &Resp{-1, nil}
}
