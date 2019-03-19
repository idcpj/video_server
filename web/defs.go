package main

const TEMPLATE_PATH = "./templates/"

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	Reqbody string `json:"reqbody"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecongized   = Err{Error: "api not reconginzed ,bad request", ErrorCode: "001"}
	ErrorRequestBodyParseFailed = Err{Error: "request body is not correct", ErrorCode: "002"}
	ErrorInternalFaults         = Err{Error: "internal service error", ErrorCode: "003"}
)
