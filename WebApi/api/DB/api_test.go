package DB

import (
	"testing"
	"log"
)
var err error
func init() {

}

func TestMain(m *testing.M){

	m.Run()
}

func TestListAllComents(t *testing.T) {
	AddNewComments("1111112","niubi",10001)
	res,err:=ListAllComents("1111111")
	if err!=nil{
		t.Errorf("List %v",err)
	}
	log.Println(res)
	dbConn.Exec("truncate comments")
}
func testListVideos(t *testing.T) {
	res,err:=AddNewVideo(10001,"fenwujiutian")
	if err!=nil{
		t.Errorf("Add:%v",err)
	}
	log.Println(*res)
	res1,err1:=ListAllVideos(10001)
	if err1!=nil{
		t.Errorf("Add:%v",err1)
	}
	log.Println(res1)
	res2,err2:=ListAllVideos(10001)
	if err2!=nil{
		t.Errorf("Add:%v",err2)
	}
	log.Println(res2[0])
	err=DelVideo(res2[0].Vid)
	if err !=nil{
		t.Errorf("DEL%v",err)
	}
}
func testGetCredential(t *testing.T){
	err=AddUserCredential("hades","199837","3211@qq.com")
	log.Println("add test")
	if err!=nil{
		t.Errorf("Add :%v",err)
	}
	_,err=GetCredential("hades")
	log.Println("GET test")
	if err!=nil{
		t.Errorf("GET :%v",err)
	}
	log.Println("DEL test")
	err=DelUser("hades","199837")
	if err!=nil{
		t.Errorf("Del:%v",err)
	}
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
}
func clearTables(){


	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}
