package middleware

import (
	"github.com/gin-gonic/gin"
	"dal/redis"
)
var login=map[string] bool{
    "/api/user/updateInfo":true,
}

func sendLogin(c *gin.Context,code int16,msg string)  {
	c.JSON(200,gin.H{
		 "code":code,
		 "msg":msg,
	})
	c.Abort();
}
func LoginAuth()  gin.HandlerFunc{
    return func(c *gin.Context) {
		if _,ok:=login[c.Request.RequestURI];ok {
           token:=c.GetHeader("token");
			if token=="" {
				sendLogin(c,401,"请先登录!")
				return
			}
			result,err:=redis.GetHash("token",token)
			if result==nil {
				sendLogin(c,401,"请先登录!")
				return
			}
			if err!=nil{
				sendLogin(c,500,"系统错误!")
				return
			}

		}
    	c.Next();
	}
}