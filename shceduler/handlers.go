package main

import (
	"net/http"
	"video_server/shceduler/dbops"

	"github.com/julienschmidt/httprouter"
)

// curl http://127.0.0.1:9001/video-delete-record/ddd
func VidDelRecHander(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, http.StatusBadRequest, "video id should not be empty")
		return
	}
	if e := dbops.AddVideoDeleionReocrd(vid); e != nil {
		sendResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	sendResponse(w, http.StatusOK, "")
	return

}
