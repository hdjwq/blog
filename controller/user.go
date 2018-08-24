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
	 if err!=nil {
        result.Code=500;
        result.Data=err;
	 }else {
		 logic.Login(account,result)
	 }
     c.Set("result",result)
}
func Register(c *gin.Context)  {
	result:=&modal.ResultData{}
	account:=&modal.Account{}
	err:=c.ShouldBind(account)
	if err!=nil {
		result.Code=500;
		result.Data=err;
	}else {
		logic.Register(account,result)
	}
	c.Set("result",result)
}

func UpdateInfo(c *gin.Context)  {
	 result:=&modal.ResultData{}
     userInfo:=&modal.UpdateUserInfo{}
     err:=c.ShouldBind(userInfo);
	 if err!=nil {
		 result.Code=500;
		 result.Data=err;
	 }else {
         logic.UpdateInfo(userInfo,result)
	 }
	c.Set("result",result)
}
