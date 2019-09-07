package DB

import (
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
var (
	dbConn *sql.DB
)
func init(){
	cfg,err:=ini.Load("E:/Golang/src/WebApi/api/DB/config.ini")
	if(err!=nil){
		log.Println(err)
		panic("Failed to load configurations")

	}
	section:=cfg.Section("mysql")
	addr:=fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",section.Key("username").String(),section.Key("password").String(),section.Key("addr").String(),section.Key("port").String(),section.Key("database").String())
	dbConn,err=sql.Open("mysql",addr)
	if(err!=nil){
		panic(err)
	}
	log.Println("initialization")
}
