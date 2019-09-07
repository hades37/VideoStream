package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandlerFunc()*httprouter.Router{
	route:=httprouter.New()
	route.GET("/video/:vid",StreamHandler)
	route.POST("video/:vid",UploadHandler)
	return route
}

func main(){
	route:=HandlerFunc()
	http.ListenAndServe(":8088",route)
}


