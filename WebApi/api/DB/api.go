package DB

import (
	"WebApi/api/defs"
	"WebApi/api/utils"
	"time"
)
func AddUserCredential(username ,password,emil string)error{
	conn,err:=dbConn.Prepare("insert into users(username,password,email) values(?,?,?)")
	if(err!=nil){
		return  err
	}
	defer conn.Close()
	conn.Exec(username,password,emil)
	return nil
}

func GetCredential(username string)(string,error){
	conn,err:=dbConn.Prepare("select uid from users where username=?")
	if err!=nil{
		return "",err
	}
	var res string
	err=conn.QueryRow(username).Scan(&res)
	if err!=nil{
		return "", err
	}
	defer conn.Close()
	return "Exist",nil
}

func DelUser(username,password string)error{
	conn,err:=dbConn.Prepare("delete from users where username=? and password=?")
	if(err!=nil){
		return err
	}
	defer conn.Close()
	_,err=conn.Exec(username,password)
	if err!=nil{
		return err
	}
	return nil
}

func EditUser(username,change string,staus int)error{
	switch staus {
	case 1:
		conn,err:=dbConn.Prepare("update from users where username=? set password=?")
		if err!=nil{
			return err
		}
		_,err=conn.Exec(username,change)
		if err!=nil{
			return err
		}

	case 2:
		conn,err:=dbConn.Prepare("update from users where username=? set email=?")
		if err!=nil{
			return err
		}
		_,err=conn.Exec(username,change)
		if err!=nil{
			return err
		}
	}

	return nil
}

func AddNewVideo(uid int,name string)(*defs.VideoInfo,error){
	vid,err:=utils.NewUUid()
	if err!=nil{
		return nil, err
	}
	t:=time.Now()
	ctime:=t.Format("Jan 02 2006,15:04:05")
	conn,er:=dbConn.Prepare("insert into video_info(vid,uid,name,time) values(?,?,?,?)")
	defer conn.Close()
	if er!=nil{
		return nil,er
	}
	_,err=conn.Exec(vid,uid,name,ctime)
	if err!=nil{
		return nil,err
	}
	return &defs.VideoInfo{
		Vid:   vid,
		Uid:   uid,
		Name:  name,
		Ctime: ctime,
	},nil
}

func ListVideos(vid string)(*defs.VideoInfo,error){
	conn,err:=dbConn.Prepare("select  uid,name,time from video_info where vid=?")
	var res *defs.VideoInfo
	res.Vid=vid
	if err!=nil{
		return nil,err
	}
	defer conn.Close()
	err=conn.QueryRow(vid).Scan(res.Uid,res.Name,res.Ctime)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

func DelVideo(vid string)error{
	conn,err:=dbConn.Prepare("delete from video_info where vid=?")
	if err!=nil{
		return err
	}
	defer conn.Close()
	_,err=conn.Exec(vid)
	if err!=nil{
		return err
	}
	return nil
}
func ListAllVideos(uid int)([]defs.VideoInfo,error){
	conn,err:=dbConn.Prepare("select vid,name,time from video_info where uid=?")
	if err!=nil{
		return nil,err
	}
	var res []defs.VideoInfo
	var temp defs.VideoInfo
	Rows,errs:=conn.Query(uid)
	if errs!=nil{
		return nil,errs
	}
	temp.Uid=uid
	for Rows.Next(){
		Rows.Scan(&temp.Vid,&temp.Name,&temp.Ctime)
		res= append(res, temp)
	}
	return res,nil
}

func AddNewComments(vid ,content string,uid int)error{
	id,err:=utils.NewUUid()
	if err!=nil{
		return err
	}
	conn,errs:=dbConn.Prepare("insert into comments(cid,vid,uid,content)values (?,?,?,?)")
	if errs !=nil{
		return errs
	}
	_,err=conn.Exec(id,vid,uid,content)
	if err!=nil{
		return err
	}
	return nil
}

func ListAllComents(vid string)([]defs.Comment,error){
	conn,err:=dbConn.Prepare("select cid,uid,content,time from comments where vid=?")
	if err!=nil{
		return nil,err
	}
	var res []defs.Comment
	var temp defs.Comment
	Rows,errs:=conn.Query(vid)
	if errs!=nil{
		return nil,errs
	}
	temp.Vid=vid
	for Rows.Next(){
		Rows.Scan(&temp.Cid,&temp.Uid,&temp.Content,&temp.Time)
		res= append(res, temp)
	}
	return res,nil
}