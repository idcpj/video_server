package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	io.WriteString(w,"hello"+name)
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user_name :=p.ByName("user_name")
	io.WriteString(w,"user_name : "+user_name)
}