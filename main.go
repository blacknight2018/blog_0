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
		article.GET("", handler.QueryArticle)
		article.POST("", handler.InsertArticle)
		article.PUT("/:id", handler.ChangeArticle)
		article.DELETE("/:id", handler.DeleteArticle)
		article.GET("/:id/detail", handler.QueryArticleDetail)
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
	comment := r.Group("/comment", handler.RequestMiddle, handler.Except)
	{
		comment.POST("", handler.InsertComment)

		comment.GET(("/:article_id"), handler.QueryComment)
	}
	r.Run(":80")
}
