package main

import (
	"html/template"
	"net/http"

	"github.com/lunny/log"

	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := &HomePage{Name: "cpj"}
	t, e := template.ParseFiles(TEMPLATE_PATH + "/home.html")
	if e != nil {
		log.Printf("Parsing template home.html error:%s", e)
	}
	t.Execute(w, p)

}
