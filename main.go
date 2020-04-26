package main

import (
	_ "blog_0/docs"
	"blog_0/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := gin.Default()

	//需要先登录
	article := r.Group("/article", handler.RequestMiddle, handler.Except)
	{
		article.GET("", handler.QueryArticle)
		article.POST("", handler.CheckLoginStatus, handler.InsertArticle)
		article.PUT("/:id", handler.CheckLoginStatus, handler.ChangeArticle)

		article.DELETE("/:id", handler.CheckLoginStatus, handler.DeleteArticle)

		article.GET("/:id/detail", handler.QueryArticleDetail)
	}

	//需要先登录
	resource := r.Group("/file", handler.RequestMiddle, handler.Except)
	{
		resource.POST("", handler.CheckLoginStatus, handler.InsertSingleFileUpload)
		resource.GET("/:id", handler.QueryFile)
	}
	user := r.Group("/user", handler.RequestMiddle, handler.Except)
	{
		user.POST("/add", handler.CheckLoginStatus, handler.InsertUser)

		user.GET("/logout", handler.DeleteUserLogout)
		//Set Session
		user.POST("/login", handler.InsertUserLogin)

		//查询用户信息用登录
		user.GET("", handler.CheckLoginStatus, handler.QueryUser)
	}
	comment := r.Group("/comment", handler.RequestMiddle, handler.Except)
	{
		comment.POST("", handler.CheckLoginStatus, handler.InsertComment)

		comment.GET(("/:article_id"), handler.QueryComment)
	}
	//解决CORS问题
	r.OPTIONS("/*all", handler.RequestMiddle, handler.Except)

	//文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
