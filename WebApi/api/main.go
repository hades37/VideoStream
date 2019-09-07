package main

import (
	R"WebApi/api/handler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type midwareHandler struct {
	r *httprouter.Router
}

func (m midwareHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	//validateUserSession(r)
	m.r.ServeHTTP(w,r)
}
func NewMidwareHandler(r *httprouter.Router)http.Handler{
	m:=midwareHandler{}
	m.r=r
	return m
}



func HandlerFunc()*httprouter.Router{
	route:=httprouter.New()
	route.POST("/user",R.CreateUser)
	route.PUT("/user/:username",R.EditUser)
	return route
}
func main(){
	log.Println("Server test Running on 8088")
	route:=HandlerFunc()
	l:=NewMidwareHandler(route)
	http.ListenAndServe(":8088",l)
}