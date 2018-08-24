package db

import (
	"com.blog/modal"
)

func FindUser(accountNumber string)(account *modal.AccountResult,err error) {
	 account=&modal.AccountResult{}
     sqlQuery:=find("*","account","account_number=?")
	 err=DB.Get(account,sqlQuery,accountNumber)
     return
}

func FindUserInfo(id int64)(userInfo *modal.UserInfo,err error){
	userInfo=&modal.UserInfo{}
	sqlQuery:=find("*","userinfo","id=?")
	err=DB.Get(userInfo,sqlQuery,id)
	return
}
func FindHasUser(accountNumber string) (num int,err error) {
     num=0;
     sqlQuery:=has("*","account","account_number=?")
     err=DB.Get(&num,sqlQuery,accountNumber)
     return
}
func Register(account *modal.Account)(userInfo *modal.UserInfo,err error){
	 sqlQuery:=insert("account","account_number,password","?,?");
	 tx,err:=DB.Begin()
	 if err!=nil {
		return
	 }
	result,err:=tx.Exec(sqlQuery,account.AccountNumber,account.Password)
	if err!=nil {
		tx.Rollback();
		return
	}
	id,err:=result.LastInsertId();
	if err!=nil {
		return
	}
	userInfo=&modal.UserInfo{};
	userInfo.Sex="3";
	userInfo.UserId=id;
	sqlQuery=insert("userinfo","id,username,sex,introduce","?,?,?,?");
	result,err=tx.Exec(sqlQuery,userInfo.UserId,"",userInfo.Sex,"");
	if err!=nil {
		tx.Rollback();
		return
	}
	tx.Commit();
	return
}

func UpdateInfo(userInfo *modal.UpdateUserInfo) error {
	 sqlQuery:=update("userinfo","username=?,sex=?,introduce=?","id=?");
	 _,err:=DB.Exec(sqlQuery,userInfo.Username,userInfo.Sex,userInfo.Introduce,userInfo.UserId);
	 return err
}