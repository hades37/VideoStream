package Router

import (
	"WebApi/api/DB"
	"WebApi/api/defs"
	"WebApi/api/errors"
	"WebApi/api/session"
	"github.com/julienschmidt/httprouter"
	"io"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(W http.ResponseWriter,R *http.Request,p httprouter.Params){
	res,_:=ioutil.ReadAll(R.Body)
	user:=&defs.UserCredential{}
	if err:=json.Unmarshal(res,user);err!=nil{
		errors.SendErrorResponse(W,defs.ErrorRequestBodyParseFailed)
		return
	}
	if err:=DB.AddUserCredential(user.Username,user.PassWord,user.Email);err!=nil{
		errors.SendErrorResponse(W,defs.ErrorDBError)
		return
	}
	id:=session.GennerSessionId(user.Username)
	su:=&defs.SignedUp{
		Success: true,
		Sid:     id,
	}
	if resp,err:=json.Marshal(su);err!=nil{
		errors.SendErrorResponse(W,defs.ErrorHttpInternal)
		return
	}else {
		errors.SendNormalResponse(W,201,string(resp))
	}

}
func GetUser(W http.ResponseWriter,R *http.Request,p httprouter.Params){


}
func EditUser(W http.ResponseWriter,R *http.Request,p httprouter.Params){
	io.WriteString(W,p.ByName("username"))
}
func DelUser(W http.ResponseWriter,R*http.Request){

}

func ListVideos(W http.ResponseWriter,R * http.Request){

}
func GetVideo(){

}
func AddVied(){

}

func DelVideo(){

}

func AddComments(){

}

func DelComments(){

}


