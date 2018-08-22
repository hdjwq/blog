package middleware

import (
	"github.com/gin-gonic/gin"
	"model"
	"gopkg.in/go-playground/validator.v8"
)
var errMap=map[string] string {
	"required":"为必填项",
}
//验证为格式错误
func validatorErr(result *model.Result) gin.H{
     errs:=result.Data;
	  switch errs.(type) {
	    case validator.ValidationErrors:
	    	 err:=errs.(validator.ValidationErrors);
			 for _,item:=range err {
				 name:=item.Name;
				 field:=item.ActualTag;
				 msg:=name;
				 if _,ok:=errMap[field];ok{
					 msg+=errMap[field]
				 }else {
					 msg+="格式错误!"
				 }
				 return gin.H{
					 "code":1002,
					 "msg":msg,
				 }
			 }
	    default:
         return gin.H{
         	"code":500,
         	"msg":"系统错误!",
		 }
	}
	return gin.H{
		"code":500,
		"msg":"系统错误!",
	}
}
//返回数据中间件
func Result() gin.HandlerFunc{
   return func(c *gin.Context) {
   	       c.Next();
	       result,ext:=c.Get("result");
	       if ext {
		      resultObj:=result.(*model.Result)
			   if resultObj.Code==200 {
				  c.JSON(200,resultObj)
			   }else if resultObj.Code==1002 {
			   	  err:=resultObj.Data.(error)
				  c.JSON(200,gin.H{
				  	  "code":1002,
				  	   "msg":err.Error(),
				  })
			   }else if resultObj.Code==500 {
				   c.JSON(200,validatorErr(resultObj))
			   }
	       }
   }
}
