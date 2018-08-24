package logic

import (
	"com.blog/modal"
	"com.blog/dal/db"
	"database/sql"
	"com.blog/uitl"
	"strconv"
)

func Login(account *modal.Account,result *modal.ResultData)  {
	accountResult,err:=db.FindUser(account.AccountNumber)
	if err!=nil {
		if err==sql.ErrNoRows {
			uitl.MakeLogicErr(result,"账号或者密码错误!")
			return
		}
        result.Code=500
        return
	}
	if accountResult.Password!=account.Password {
		uitl.MakeLogicErr(result,"账号或者密码错误!")
		return
	}
	userInfo,err:=db.FindUserInfo(accountResult.Id)
	if err!=nil {
		result.Code=500
		return
	}
	userInfo.Token=uitl.MakeToken(userInfo.UserId);
    err=hasToken(userInfo.UserId);
	if err!=nil {
		result.Code=500
		return
	}
	err=setToken(userInfo.Token,userInfo.UserId);
	if err!=nil {
		result.Code=500
		return
	}
	result.Code=200;
	result.Data=userInfo;
}
func Register(account *modal.Account,result *modal.ResultData)  {
	 num,err:=db.FindHasUser(account.AccountNumber);
	 if err!=nil{
		result.Code=500
		return
	 }
	 if num!=0{
		 uitl.MakeLogicErr(result,"账号已经存在!")
		 return
	 }
     userInfo,err:=db.Register(account);
	 if err!=nil {
		 result.Code=500
		 return
	 }
	 userInfo.Token=uitl.MakeToken(userInfo.UserId);
	 err=setToken(userInfo.Token,userInfo.UserId);
	 if err!=nil {
		result.Code=500
		return
	 }
	 result.Code=200;
	 result.Data=userInfo;
}
func UpdateInfo(userInfo *modal.UpdateUserInfo,result *modal.ResultData){
	sex,err:=strconv.Atoi(userInfo.Sex);
	if err!=nil {
		uitl.MakeLogicErr(result,"sex格式错误!")
		return
	}
	if sex<0||sex>2 {
		uitl.MakeLogicErr(result,"sex格式错误!")
		return
	}
    err=db.UpdateInfo(userInfo);
	if err!=nil {
        result.Code=500;
        return
	}
	result.Code=200;
	result.Data=userInfo;
}
