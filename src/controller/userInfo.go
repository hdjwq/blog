package controller

import (
	"github.com/gin-gonic/gin"
	"logic"
	"model"
	"model/user"
)
//登录
func Login(c *gin.Context)  {
	 result:=&model.Result{}
	 login:=&user.LoginOrRegister{}
	 if err:=bindObj(c,result,login);err!=nil{
	 	return
	 }
	 logic.Login(login,result);
	 c.Set("result",result);
}
//注册
func Register(c *gin.Context)  {
	result:=&model.Result{}
	register:=&user.LoginOrRegister{}
	if err:=bindObj(c,result,register);err!=nil{
		return
	}
	logic.Register(register,result);
	c.Set("result",result);
}
//用户信息更新
func UpdateInfo(c *gin.Context)  {
	result:=&model.Result{}
	updateInfo:=&user.UpdateInfo{}
	if err:=bindObj(c,result,updateInfo);err!=nil{
		return
	}
    logic.UpdateInfo(updateInfo,result)
}
