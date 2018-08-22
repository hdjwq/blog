package db

import (
	"model/user"
	"database/sql"
	"fmt"
)
//登录
func Login(accountNumber string) (resultInfo *user.LoginResult,err error) {
	 resultInfo=&user.LoginResult{}
	 sqlQuery:=find("id,password","account","account_number=?");
	 err=DB.Get(resultInfo,sqlQuery,accountNumber)
	 if err!=nil {
	    return
	 }
	 return
}
//获取用户信息
func GetUserInfo(id int64)(result *user.UserInfoResult,err error){
	result=&user.UserInfoResult{}
	sqlQuery:=find("*","userinfo","id=?")
	err=DB.Get(result,sqlQuery,id)
	return
}
//注册
func Register(register *user.LoginOrRegister)(resultInfo *user.UserInfoResult,err error){
	 var tx *sql.Tx
	 var result sql.Result;
	 resultInfo=&user.UserInfoResult{};
	 tx,err=DB.Begin()
	 if err!=nil {
		return
	 }
	 sqlQurey:=insert("account_number,password","?,?","account")
	 result,err=tx.Exec(sqlQurey,register.AccountNumber,register.PassWord);
	 if err!=nil {
        tx.Rollback()
        return
	 }
	 var id int64
	 id,err=result.LastInsertId()
	if err!=nil {
		tx.Rollback();
		return
	}
	resultInfo.UserId=id;
	resultInfo.Sex=3;
	sqlQurey=insert("id,username,sex,introduce","?,?,?,?","userinfo")
	result,err=tx.Exec(sqlQurey,id,"",3,"")
	if err!=nil {
		tx.Rollback();
		return
	}
	tx.Commit();
	return
}
//更新数据
func UpdateInfo(updateInfo *user.UpdateInfo) error {
     sqlQuery:=update("userinfo","username=?,sex=?,introduce=?","id=?")
     fmt.Println(updateInfo.Id)
     result,err:=DB.Exec(sqlQuery,updateInfo.UserName,updateInfo.Sex,updateInfo.Introduce,updateInfo.Id)
     fmt.Println(result.RowsAffected())
     return err
}

