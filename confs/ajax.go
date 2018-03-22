package confs

//Code reponse code value
type Code int

const (
	//SUCCESS code
	SUCCESS Code = 0
)

//Resp the REST response model
type Resp struct {
	Code       Code        `json:"code"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}
