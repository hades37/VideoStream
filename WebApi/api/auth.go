package main

import (
	"WebApi/api/defs"
	"WebApi/api/errors"
	"WebApi/api/session"
	"net/http"
)


var HEADER_SESSION="X-Session-Id"
var HEADER_USER="X-User-Name"
func validateUserSession(r *http.Request)bool{
	sid:=r.Header.Get(HEADER_SESSION)
	if len(sid)==0{
		return false
	}
	uname,ok:=session.IsSessionExpired(sid)
	if ok{
		return false
	}
	r.Header.Add(HEADER_USER,uname)
	return true
}

func validUser(w http.ResponseWriter,r *http.Request)bool{
	uname:=r.Header.Get(HEADER_USER)
	if len(uname)==0{
		errors.SendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return false
	}
return true
}