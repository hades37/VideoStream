package errors

import (
	"WebApi/api/defs"
	"encoding/json"
	"io"
	"net/http"
)

func SendErrorResponse(W http.ResponseWriter, err defs.ErrorResponse){
	W.WriteHeader(err.HttpSC)
	reStr,_:=json.Marshal(&err.Error)
	io.WriteString(W,string(reStr))
}

func SendNormalResponse(W http.ResponseWriter,cd int,msg string){
	W.WriteHeader(cd)
	io.WriteString(W,msg)
}