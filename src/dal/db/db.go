package db

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	 "fmt"
)
var DB *sqlx.DB

func Init(addr string)  {
	 var err error
	 DB,err=sqlx.Connect("mysql",addr)
	if err!=nil {
		panic(fmt.Sprintf("连接数据库错误:%v",err))
	}
}