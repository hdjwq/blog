package controller

import (
	"github.com/gin-gonic/gin"
	"com.blog/logic"
	"com.blog/modal"
)
func Login(c *gin.Context)  {
	 result:=&modal.ResultData{}
     account:=&modal.Account{}
     err:=c.ShouldBind(account)
	 defer setResult(c,result)
	 if err!=nil {
        result.Code=500;
        result.Data=err;
        return
	 }
	 logic.Login(account,result)
	 return
}
func Register(c *gin.Context)  {
	result:=&modal.ResultData{}
	account:=&modal.Account{}
	err:=c.ShouldBind(account)
	defer setResult(c,result)
	if err!=nil {
		result.Code=500;
		result.Data=err;
		return
	}
	logic.Register(account,result)
	return
}

func UpdateInfo(c *gin.Context)  {
	 result:=&modal.ResultData{}
     userInfo:=&modal.UpdateUserInfo{}
     err:=c.ShouldBind(userInfo);
	 defer setResult(c,result)
	 if err!=nil {
		result.Code=500;
		result.Data=err;
		return
	 }
	userId,_:=c.Get("userId");
    userNum,err:=getUserId(userId);
	if err!=nil {
		result.Code=500;
		return
	}
	userInfo.UserId=userNum;
	logic.UpdateInfo(userInfo,result)
	return
}
