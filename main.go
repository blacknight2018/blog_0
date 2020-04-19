package main

import (
	"blog_0/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	article := r.Group("/article", handler.Except, handler.RequestMiddle)
	{
		article.GET("/", handler.Query)
		article.POST("/", handler.Insert)
		article.PUT("/:id", handler.Change)
		article.DELETE("/:id", handler.Delete)
		article.GET("/:id/detail", handler.QueryDetail)
	}
	resource := r.Group("/file", handler.Except, handler.RequestMiddle)
	{
		resource.POST("/", handler.SingleFileUpload)
		resource.GET("/:id", handler.GetFile)
	}
	r.Run(":80")
}
