package main

import (
	"net/http"

	"os"

	"time"

	"log"

	"github.com/julienschmidt/httprouter"
)

//link http://127.0.0.1:9000/videos/test1
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	vid := p.ByName("vid-id")
	filepath := VIDEO_DIR + vid

	video, e := os.Open(filepath)
	defer video.Close()
	if e != nil {
		log.Printf("filepath :%s   error: ", filepath, e)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	//把文件放入响应中
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
