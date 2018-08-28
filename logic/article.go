package logic

import (
	"com.blog/modal"
	"com.blog/dal/db"
	"com.blog/uitl"
	"database/sql"
)

func AddArticle(data *modal.Article,result *modal.ResultData) {
  if data.Synopsis=="" {
  	    runeSlice:=[] rune(data.Content)
  	    lens:=len(runeSlice);
	    if lens>100 {
		  lens=100
	    }
		strSlice:=runeSlice[:lens]
	  data.Synopsis=string(strSlice);
  }
	data.Status=1;
  err:=db.AddArticle(data);
  if err!=nil {
    result.Code=500;
    return
  }
  result.Code=200;
}
func DeleteArticle(data *modal.ArticleIdData,result *modal.ResultData)  {
	 info:=&modal.ArticleInfo{}
	 err:=db.FindArticleInfo(info,data.ArticleId)
	 if err!=nil {
		 if err==sql.ErrNoRows {
			 uitl.MakeLogicErr(result,"该文章不存在!");
			 return
		 }
		 result.Code=500
		 return
	 }
	 if info.Status==0 {
		uitl.MakeLogicErr(result,"该文章已经被删除!");
		return
	 }
     err=db.DeleteArticle(data.ArticleId)
	 if err!=nil {
		result.Code=500
		return
	 }
	result.Code=200
	return
}
func ArticleInfo(data *modal.ArticleIdData,result *modal.ResultData)  {
     resultData:=&modal.ArticleInfo{}
     err:=db.FindArticleInfo(resultData,data.ArticleId)
	 if err!=nil {
		 if err==sql.ErrNoRows {
			 uitl.MakeLogicErr(result,"该文章不存在!");
			 return
		 }
		result.Code=500;
		return
	 }
     err=db.FindUserName(resultData,resultData.UserId);
	 if err!=nil {
		result.Code=500;
		return
	 }
	 if resultData.Status==0 {
		 uitl.MakeLogicErr(result,"该文章已经被删除!");
		 return
	 }
	 result.Code=200;
	 result.Data=resultData;
	 return
}
