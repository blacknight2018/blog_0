package main

import (
	_ "blog_0/docs"
	"blog_0/handler"
	"blog_0/handler/articleSer"
	"blog_0/handler/commentSer"
	"blog_0/handler/resourceSer"
	"blog_0/handler/userSer"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "net/http/pprof"
)

func main() {

	r := gin.Default()
	pprof.Register(r)
	//需要先登录

	article := r.Group("/article", handler.RequestMiddle, handler.Except)
	{
		article.GET("", articleSer.QueryArticle)
		article.POST("", handler.CheckLoginStatus, articleSer.InsertArticle)
		article.PUT("/:id", handler.CheckLoginStatus, articleSer.ChangeArticle)

		article.DELETE("/:id", handler.CheckLoginStatus, articleSer.DeleteArticle)

		article.GET("/:id/detail", articleSer.QueryArticleDetail)
	}

	//需要先登录
	resource := r.Group("/file", handler.RequestMiddle, handler.Except)
	{
		resource.POST("", handler.CheckLoginStatus, resourceSer.InsertSingleFileUpload)
		resource.GET("/:id", resourceSer.QueryFile)
	}
	user := r.Group("/user", handler.RequestMiddle, handler.Except)
	{
		user.POST("/add", userSer.InsertUser)

		user.DELETE("/logout", userSer.DeleteUserLogout)
		//Set Session
		user.POST("/login", userSer.InsertUserLogin)

		//查询用户信息用登录
		user.GET("", handler.CheckLoginStatus, userSer.QueryUser)
	}
	comment := r.Group("/comment", handler.RequestMiddle, handler.Except)
	{
		comment.POST("", handler.CheckLoginStatus, commentSer.InsertComment)

		comment.GET(("/:article_id"), commentSer.QueryComment)
	}
	//解决CORS问题
	r.OPTIONS("/*all", handler.RequestMiddle, handler.Except)

	r.GET("", articleSer.QueryArticle)
	//文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
