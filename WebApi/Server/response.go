package main

import (
	"io"
	"net/http"
)

func erroStatuResponse(W http.ResponseWriter,code int, MSg string){
	W.WriteHeader(code)
	io.WriteString(W,MSg)
}