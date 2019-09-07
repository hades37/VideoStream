package DB

import ("strconv"
	"WebApi/api/defs"
	"sync"
)

func AddNewSession(sid string,ttl int64,name string)error{
	ttlString:=strconv.FormatInt(ttl,10)
	conn,err:=dbConn.Prepare("insert into sessions(sid,ttl,username)values(?,?,?)")
	if err!=nil{
		return err
	}
	defer conn.Close()
	_,err=conn.Exec(sid,ttlString,name)
	if err!=nil{
		return err
	}
	return nil
}

func GetSession(sid string)(*defs.Session,error){
	var session=&defs.Session{}
	conn,err:=dbConn.Prepare("select username,ttl from sessions where sid=?")
	if err!=nil{
		return nil,err
	}
	defer conn.Close()
	var ttl string
	conn.QueryRow(sid).Scan(session.Name,&ttl)
	session.TTL,err=strconv.ParseInt(ttl,10,64)
	if err!=nil{
		return nil,err
	}
	return session,nil
}

func ListAllSessions()(*sync.Map,error){
	var Map=&sync.Map{}
	conn,err:=dbConn.Prepare("select  username ,sid ,ttl from sessions")
	if err!=nil{
		return nil,err
	}
	defer conn.Close()
	rows,errs:=conn.Query()
	if errs!=nil{
		return nil,errs
	}
	var Recored=&defs.Session{}
	var sid ,ttl string
	for rows.Next(){
		rows.Scan(Recored.Name,&sid,&ttl)
		Recored.TTL,err=strconv.ParseInt(ttl,10,64)
		Map.Store(sid,Recored)
	}
	return Map,nil
}

func DelSession(sid string)error{
	conn,err:=dbConn.Prepare("delete from sessions where sid=?")
	if err!=nil{
		return err
	}
	defer conn.Close()
	_,err=conn.Exec(sid)
	if err!=nil{
		return err
	}
	return nil
}