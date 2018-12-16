package main

import (
	"net/http"
	"video_server/api/session"
)

const (
	//X 开头的是用户自定义头
	HEADER_FIELD_SESSION = "X-Session-Id"
	HEADER_FIELD_UNAME   = "X-User-Name"
)

//session 校验
func ValidateUserSession(r *http.Request) bool {
	//获取 请求头中的 HEADER_FIELD_SESSION 值
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	name, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, name)
	return true

}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w)
		return false
	}

}
