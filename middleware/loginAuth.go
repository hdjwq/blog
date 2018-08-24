package middleware

import (
	"github.com/gin-gonic/gin"
	"com.blog/dal/redis"
)
var loginMap=map[string] bool{
	"/api/user/updateInfo":true,
}

func toLogin(c *gin.Context)  {
	c.JSON(200,gin.H{
		"code":401,
		"msg":"请先登录",
	})
	c.Abort()
}
func LoginAuth()  gin.HandlerFunc{
	return func(c *gin.Context) {
		   url:=c.Request.RequestURI;
		   if _,ok:=loginMap[url];ok {
			  token:=c.GetHeader("token");
			   if token=="" {
				   toLogin(c)
				   return
			   }
			  result,err:=redis.HGet("token",token)
			   if err!=nil||result==nil {
				   toLogin(c)
				   return
			   }
		  }
		  c.Next();
	}
}
