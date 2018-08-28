package controller

import (
	"github.com/gin-gonic/gin"
	"com.blog/modal"
	"com.blog/logic"
)

func AddArticle(c *gin.Context)  {
	 article:=&modal.Article{};
	 result:=&modal.ResultData{}
	 err:=c.ShouldBind(article);
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
	 article.UserId=userNum;
	 logic.AddArticle(article,result)
	 return
}
func DeleteArticle(c *gin.Context)  {
	deleteData:=&modal.ArticleIdData{};
	result:=&modal.ResultData{}
	err:=c.ShouldBind(deleteData);
	defer setResult(c,result)
	if err!=nil {
		result.Code=500;
		result.Data=err;
		return
	}
    logic.DeleteArticle(deleteData,result)
	return
}
func ArticleInfo(c *gin.Context)  {
	articleData:=&modal.ArticleIdData{};
	result:=&modal.ResultData{};
	err:=c.ShouldBind(articleData);
	defer setResult(c,result)
	if err!=nil {
		result.Code=500;
		result.Data=err;
		return
	}
	logic.ArticleInfo(articleData,result)
	return
}