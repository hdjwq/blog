package db

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var DB *sqlx.DB
func Init(addr string)  {
	 var err error
	 DB,err=sqlx.Connect("mysql",addr)
	 if err!=nil {
		panic(err)
	 }

}