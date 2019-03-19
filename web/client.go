package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error
	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		//build quest header
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Print(err)
			return
		}
		normalReqponse(w, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.Reqbody)))

		//build quest header
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Print(err)
			return
		}
		normalReqponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", b.Url, nil)

		//build quest header
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Print(err)
			return
		}
		normalReqponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad api request")
	}

}

func normalReqponse(w http.ResponseWriter, w1 *http.Response) {
	res, e := ioutil.ReadAll(w1.Body)
	if e != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))

	}

	w.WriteHeader(w1.StatusCode)
	io.WriteString(w, string(res))
}
