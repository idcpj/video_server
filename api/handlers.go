package main

import (
	"io"
	"io/ioutil"
	"net/http"

	"encoding/json"
	"video_server/api/defs"

	"video_server/api/dbops"

	"video_server/api/session"

	"github.com/julienschmidt/httprouter"
)

//@link curl -d '{"user_name":"idcpj","pwd":"12345"}' -H 'Content-Type: application/json' -X POST http://127.0.0.1:8000/user
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if e := json.Unmarshal(res, ubody); e != nil {
		sendErrorResponse(w, defs.ErrorResponseBodyParseFailed)
		return
	}
	if e := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); e != nil {
		sendErrorResponse(w, defs.ErrorDbError)
		return
	}
	id := session.GenerateNewSessionId(ubody.UserName)
	su := &defs.SignedUp{Success: true, SessionId: id}
	if resp, e := json.Marshal(su); e != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user_name := p.ByName("user_name")
	io.WriteString(w, "user_name : "+user_name)
}
