package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", HomeHandler)
	router.POST("/", HomeHandler)

	router.GET("/userhome", UserHomeHandler)
	router.POST("/userhome", UserHomeHandler)

	router.POST("/api", Apihandler)

	router.POST("/upload/:vid-id", proxyUploadHandler)

	router.ServeFiles("/statics/*filepath", http.Dir(TEMPLATE_PATH))

	return router
}

func main() {

	r := RegisterHandler()
	e := http.ListenAndServe(":8000", r)
	if e != nil {
		log.Print("create server is error")
		return
	}
	log.Println("create server is success ,the port is 8000")

}
