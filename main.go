package main

import (
	"github.com/gin-gonic/gin"
    "com.blog/dal/db"
    "com.blog/controller"
    "com.blog/middleware"
    "com.blog/dal/redis"
)
func main()  {
	db.Init("root:123456@tcp(localhost:3306)/blog?charset=utf8&parseTime=true");
	redis.Init("localhost:6379");
	r:=gin.Default();
	r.Use(middleware.LoginAuth());
	r.Use(middleware.ResultMiddleware());
	router:=r.Group("/api");
	{
		user:=router.Group("/user");
		{
			user.POST("/login",controller.Login);
			user.POST("/register",controller.Register);
			user.POST("/updateInfo",controller.UpdateInfo);
		};
		article:=router.Group("/article");
		{
			article.POST("/add",controller.AddArticle);
			article.POST("/delete",controller.DeleteArticle);
			article.GET("/info",controller.ArticleInfo)
		}
	}
	r.Run()
}
