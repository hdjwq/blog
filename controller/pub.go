package controller

import (
	"github.com/gin-gonic/gin"
	"com.blog/modal"
	"fmt"
	"strconv"
)

func setResult(c *gin.Context,result *modal.ResultData)  {
	 c.Set("result",result)
}
func getUserId(id interface{})(userNum int64,err error){
	userIdStr:=fmt.Sprintf("%v",id)
	userNum,err=strconv.ParseInt(userIdStr,10,64)
	return
}
