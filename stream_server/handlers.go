package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"os"

	"time"

	"log"

	"github.com/julienschmidt/httprouter"
)

//也可以直接用 curl 模拟
// curl --form file=@/Users/idcpj/go/src/video_server/testData/test.mov --form press=ok http://127.0.0.1:9000/upload/ggg
func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, e := template.ParseFiles(VIDEO_DIR_2 + "upload.html")
	if e != nil {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		log.Printf("upload file error :%v  filepath :%v", e, dir)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	t.Execute(w, nil)
}

//link http://127.0.0.1:9000/videos/test1
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	vid := p.ByName("vid-id")
	filePath := VIDEO_DIR_2 + vid

	video, e := os.Open(filePath)
	defer video.Close()
	if e != nil {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		log.Printf("filePath :%v   error:%v   absolute_path:%v", filePath, e, dir)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	//把文件放入响应中
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

}

//文件上传
// curl
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if e := r.ParseMultipartForm(MAX_UPLOAD_SIZE); e != nil {
		log.Println(" ParseMultipartForm is error ", e)
		sendErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	file, _, e := r.FormFile("file")
	if e != nil {
		log.Println(" FormFile is error ", e)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	data, e := ioutil.ReadAll(file)
	if e != nil {
		log.Println(" readfile is error ", e)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	vidId := p.ByName("vid-id")
	e = ioutil.WriteFile(VIDEO_DIR_2+vidId, data, 0666)
	if e != nil {
		s, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		log.Printf(" WriteFile is error:%v path_dir :%s", e, s)
		sendErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload successfully")

}
