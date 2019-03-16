package main

import (
	"net/http"

	"github.com/lunny/log"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", HomeHandler)
	router.POST("/", HomeHandler)

	//router.GET("/userhome", userHomeHandler)
	//router.POST("/userhome", userHomeHandler)
	//
	//router.POST("/api", apihandler)

	router.ServeFiles("/statics/*filepath", http.Dir(TEMPLATE_PATH))

	return router
}

func main() {

	r := RegisterHandler()
	e := http.ListenAndServe(":8000", r)
	if e != nil {
		log.Error("create server is error")
	}
	log.Println("create server is success ,the port is 8000")

}
