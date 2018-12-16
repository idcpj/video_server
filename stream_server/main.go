package main

import (
	"log"
	"net/http"

	"os"

	"github.com/julienschmidt/httprouter"
)

type middleware struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMidWare(r *httprouter.Router, cc int) *middleware {
	log.Printf("new midWare")
	m := &middleware{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("getCon ...")
	os.Exit(0)
	if !m.l.GetConn() {
		log.Println("getCon is error")
		sendErrorResponse(w, http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()

}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	return router
}

func main() {
	log.Println("init stream_server post is 9000")
	r := RegisterHandlers()
	m := NewMidWare(r, 2)
	http.ListenAndServe(":9000", m.r)
}
