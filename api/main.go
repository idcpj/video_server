package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type minddleWarehandle struct {
	r *httprouter.Router
}

func NewMiddleWareHandle(r *httprouter.Router) minddleWarehandle {
	m := minddleWarehandle{}
	m.r = r
	return m
}

//在 原始的 ServerHttp 上封装一层 ,中间件的原理
func (m minddleWarehandle) ServerHttp(w http.ResponseWriter, r *http.Request) {
	//check session
	ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisetHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisetHandlers()
	//mh := NewMiddleWareHandle(r)

	http.ListenAndServe(":8000", r)
}
