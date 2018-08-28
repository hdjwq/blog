package modal

type Article struct {
	 UserId int64 `json:"userId" db:"user_id"`
	 Title string `json:"title" db:"title" form:"title" binding:"required,min=1,max=40"`
	 Types string `json:"types" db:"types" form:"types" binding:"required,number"`
	 Content string `json:"content" db:"content" form:"content" binding:"required"`
	 Synopsis string `json:"synopsis" db:"synopsis" form:"synopsis"`
	 Status int8 `db:"status"`
}
type ArticleIdData struct {
	 ArticleId int64 `form:"articleId" binding:"required"`
}
type ArticleInfo struct {
	Article
	UserName string `json:"userName" db:"username"`
	Id int64 `json:"id" db:"id"`
}