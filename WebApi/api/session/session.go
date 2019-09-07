package session

import (
	"WebApi/api/defs"
	"WebApi/api/utils"
	"sync"
	"WebApi/api/DB"
	"time"
)


var sessionMap *sync.Map
func init(){
	sessionMap=&sync.Map{}
}

func LoadSessionFromDB(){
	s,err:=DB.ListAllSessions()
	if err!=nil{
			return
			}
	s.Range(func(key, value interface{}) bool {
		ss:=value.(*defs.Session)
		sessionMap.Store(key,ss)
		return true
	})
}

func GennerSessionId(username string)string{
	id,_:=utils.NewUUid()
	ct:=time.Now().UnixNano()/1000000
	ttl:=ct+30*60*1000
	ss:=&defs.Session{
		Name: username,
		TTL:  ttl,
	}
	sessionMap.Store(id,ss)
	DB.AddNewSession(id,ttl,username)
	return id
}

func IsSessionExpired(sid string)(string,bool){
	ss,ok:=sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano() / 1000000
		if ss.(*defs.Session).TTL < ct {
			DB.DelSession(sid)
			sessionMap.Delete(sid)
			return "", true
		}
		return ss.(defs.Session).Name,false
	}
		return "",true
}