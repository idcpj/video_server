package main

import (
	"log"
	"net/http"
	"video_server/shceduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", VidDelRecHander)
	return router
}

func main() {
	go taskrunner.Start()

	handlers := RegisterHandlers()
	e := http.ListenAndServe(":9001", handlers)
	if e != nil {
		log.Println(e)
	}
}
