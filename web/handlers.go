package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/lunny/log"

	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Print(r.URL.String())
	username, e1 := r.Cookie("username")
	sid, e2 := r.Cookie("session")

	//no login
	if e1 == http.ErrNoCookie || e2 == http.ErrNoCookie {
		p := &HomePage{Name: "cpj"}
		t, e := template.ParseFiles(TEMPLATE_PATH + "home.html")
		if e != nil {
			log.Printf("Parsing template home.html error:%s", e)
		}
		t.Execute(w, p)
		return
	}

	//login
	if len(username.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/username", http.StatusFound)
		return
	}

}
func UserHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Print(r.URL.String())
	username, e1 := r.Cookie("username")
	_, e2 := r.Cookie("session")

	//first go to this page
	if e1 == http.ErrNoCookie || e2 == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	var p UserPage
	fname := r.FormValue("username")
	if len(fname) != 0 {
		p = UserPage{Name: username.Value}
	} else {
		p = UserPage{Name: fname}
	}
	t, e3 := template.ParseFiles(TEMPLATE_PATH + "userhome.html")
	if e3 != nil {
		log.Errorf("parsing username.html error:%s", e3)
		return
	}

	t.Execute(w, p)

}

func Apihandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// not post
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecongized)
		io.WriteString(w, string(re))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if e := json.Unmarshal(res, apiBody); e != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}

	request(apiBody, w, r)
	defer r.Body.Close()

}
