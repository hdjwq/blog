package logic

import (
	"model/user"
	"dal/db"
	"model"
	"errors"
	"database/sql"
	"uitl"
	"fmt"
	"dal/redis"
)
//登录
func Login(login *user.LoginOrRegister,result *model.Result){
     loginResult,err:=db.Login(login.AccountNumber)
	if err==sql.ErrNoRows {
		result.Code=1002
		result.Data=errors.New("账号不存在!")
		return
	}
	if err!=nil {
		uitl.SysErr(result)
		return
	}
	if loginResult.PassWord!=login.PassWord {
		result.Code=1002
		result.Data=errors.New("账号或密码错误!")
		return
	}
	userInfoResult,err:=db.GetUserInfo(loginResult.Id);
	if err!=nil {
		uitl.SysErr(result)
		return
	}
	userInfoResult.Token=uitl.MakeToken(userInfoResult.UserId)
	redisResult,err:=redis.GetHash("id",userInfoResult.UserId)
	if err!=nil {
		uitl.SysErr(result)
		return
	}
	tokenStr:=fmt.Sprintf("%s",redisResult)
	err=redis.RmHash("token",tokenStr)
	if err!=nil {
		uitl.SysErr(result)
		return
	}
	err=redis.SetHash("id",userInfoResult.UserId,userInfoResult.Token)
	if err!=nil {
		uitl.SysErr(result)
		return
	}
	err=redis.SetHash("token",userInfoResult.Token,userInfoResult.UserId)
	if err!=nil {
		uitl.SysErr(result)
		return
	}
    uitl.Success(result,userInfoResult)
}
//注册
func Register(login *user.LoginOrRegister,result *model.Result)  {
	resultInfo,errs:=db.Register(login);
	if errs!=nil {
		sqlErrRegister(errs,result)
		return
	}
	resultInfo.Token=uitl.MakeToken(resultInfo.UserId);
	errs=redis.SetHash("token",resultInfo.Token,resultInfo.UserId)
	errs=redis.SetHash("id",resultInfo.UserId,resultInfo.Token)
	if errs!=nil {
		uitl.SysErr(result)
		return
	}
	uitl.Success(result,resultInfo)
}

func UpdateInfo(updateInfo *user.UpdateInfo,result *model.Result){
     err:=db.UpdateInfo(updateInfo)
     fmt.Println(err)
}