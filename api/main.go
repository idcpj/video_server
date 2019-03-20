package main

import (
	"net/http"

	"os"

	"github.com/julienschmidt/httprouter"
)

type minddleWarehandle struct {
	r *httprouter.Router
}

func NewMiddleWareHandle(r *httprouter.Router) minddleWarehandle {
	m := minddleWarehandle{r: r}
	return m
}

// !!! 覆写 ServerHttp 接口
func (m minddleWarehandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	os.Exit(0)
	ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisetHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:username", GetUserInfo)
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/comments", ShowComments)
	return router
}

func main() {
	r := RegisetHandlers()
	mh := NewMiddleWareHandle(r)

	http.ListenAndServe(":8000", mh.r)
}
