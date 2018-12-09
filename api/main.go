package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)



func RegisetHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisetHandlers()
	http.ListenAndServe(":8000", r)
}
