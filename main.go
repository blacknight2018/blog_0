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

	article := r.Group("/articles", handler.RequestMiddle, handler.Except)
	{
		article.GET("", handler.Query)
		article.POST("", handler.Insert)
		article.PUT("/:id", handler.Change)
		article.DELETE("/:id", handler.Delete)
		article.GET("/:id/detail", handler.QueryDetail)
	}
	resource := r.Group("/file", handler.RequestMiddle, handler.Except)
	{
		resource.POST("/", handler.SingleFileUpload)
		resource.GET("/:id", handler.GetFile)
	}
	user := r.Group("/user", handler.RequestMiddle, handler.Except)
	{
		user.POST("/add", handler.UserInsert)

		//Set Session
		user.POST("/login", handler.UserLogin)

		user.GET("", handler.UserQuery)
	}
	r.Run(":80")
}
