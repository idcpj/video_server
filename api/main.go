package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisetHandle() *httprouter.Router {
	router := httprouter.New()
	router.POST("/User", createUser)
	return router
}

func main() {
	r := RegisetHandle()
	http.listenAndServer(":8000", r)
}
