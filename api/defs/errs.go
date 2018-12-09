package defs

type Err struct{
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct{
	HttpSc int
	Error Err
}
var (
	ErrorResponseBodyParseFailed = ErroResponse{HttpSc:400,Error:Err{Error:"Request body is not correct",ErrorCode:"0001"}}
	ErrorNotAuthUser = ErroResponse{HttpSc:401,Error:Err{Error:"User authentication failed.",ErrorCode:"002"}}
)