package controller

import (
	"github.com/gin-gonic/gin"
	"model"
)
func bindObj(c *gin.Context,result *model.Result,data interface{}) error{
	err:=c.ShouldBind(data);
	if err!=nil {
		result.Code=500;
		result.Data=err;
		c.Set("result",result);
	}
	return err
}