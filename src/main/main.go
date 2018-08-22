package main

import (
	"github.com/gin-gonic/gin"
	"controller"
	"dal/db"
	"middleware"
	"dal/redis"
)

func main()  {
	 db.Init("root:123456@tcp(localhost:3306)/blog")
	 redis.Init("localhost:6379")
	 r:=gin.Default();
	 r.Use(middleware.LoginAuth())
	 r.Use(middleware.Result())
	 product:=r.Group("/api")
	 user:=product.Group("/user")
	 {
	 	user.POST("/login",controller.Login)
	 	user.POST("/register",controller.Register)
	 	user.POST("/updateInfo",controller.UpdateInfo)
	 }
	 r.Run(":8080")
}