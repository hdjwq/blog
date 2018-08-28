package db

import (
	"com.blog/modal"
)
func AddArticle(a *modal.Article) error {
	 sqlQuery:=insert("article","user_id,title,types,content,synopsis,status","?,?,?,?,?,?");
	 _,err:=DB.Exec(sqlQuery,a.UserId,a.Title,a.Types,a.Content,a.Synopsis,a.Status);
	 return err
}
func DeleteArticle(id int64) error {
	 sqlQuery:=update("article","status=?","id=?");
	 _,err:=DB.Exec(sqlQuery,0,id)
	 return err
}
func FindArticleInfo(data interface{},id int64)error{
	 sqlQuery:=find("*","article","id=?");
	 return DB.Get(data,sqlQuery,id);
}
func FindUserName(data interface{},id int64) error {
	 sqlQuery:=find("username","userinfo","id=?");
	return DB.Get(data,sqlQuery,id);
}