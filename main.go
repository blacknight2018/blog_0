package main

import (
	"blog_0/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func main() {

	//
	r := gin.Default()
	//Session profile

	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions(configure.SessionName, store))

	//需要先登录
	article := r.Group("/articles", handler.RequestMiddle, handler.Except)
	{
		article.GET("", handler.Query)
		article.POST("", handler.Insert)
		article.PUT("/:id", handler.Change)
		article.DELETE("/:id", handler.Delete)
		article.GET("/:id/detail", handler.QueryDetail)
	}

	//需要先登录
	resource := r.Group("/file", handler.RequestMiddle, handler.Except)
	{
		resource.POST("/", handler.SingleFileUpload)
		resource.GET("/:id", handler.GetFile)
	}
	user := r.Group("/user", handler.RequestMiddle, handler.Except)
	{
		//需要先登录
		user.POST("/add", handler.UserInsert)

		//Set Session
		user.OPTIONS("/login")
		user.POST("/login", handler.UserLogin)

		//需要先登录
		user.GET("", handler.UserQuery)
	}
	r.Run(":80")
}
