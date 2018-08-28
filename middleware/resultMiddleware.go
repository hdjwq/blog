package middleware

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"com.blog/modal"
)
var errCode=map[string] string{
	"required":"为必填项",
}
func isValidatorErr(c *gin.Context,data interface{})  {
	switch data.(type) {
	case validator.ValidationErrors:
         resultData:=data.(validator.ValidationErrors);
		for _,item:=range resultData{
			var errStr string
			var ok bool
			name:=item.Name;
			field:=item.ActualTag;
			if errStr,ok=errCode[field];!ok{
				errStr="格式错误!"
			}
			c.JSON(200,gin.H{
				"code":1002,
				"msg":name+errStr,
			})
			break
		}
	default:
        c.JSON(200,gin.H{
        	"code":500,
        	"msg":"系统错误!",
		})
	}
}
func ResultMiddleware()  gin.HandlerFunc{
	 return func(c *gin.Context) {
	 	    c.Next();
            result,exists:=c.Get("result")
		    if exists {
				resultObj:=result.(*modal.ResultData);
				switch resultObj.Code {
				 case 200:
				 	c.JSON(200,gin.H{
				 		"code":200,
				 		 "data":resultObj.Data,
				 		 "msg":"操作成功!",
					})
				 case 1002:
				 	 err:=resultObj.Data.(error)
					 c.JSON(200,gin.H{
						 "code":1002,
						 "msg":err.Error(),
					 })
				 case 500:
					isValidatorErr(c,resultObj.Data)
				}
               return
		    }
		    c.JSON(200,gin.H{
		    		 "code":404,
		    		 "msg":"没有找到!",
			})
	 }
}
