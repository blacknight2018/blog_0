package main

import (
	"blog_0/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	//s := struct {
	//	Name string
	//	Age  int
	//}{"aa", 18}
	//
	//b ,_:= setKeyValue(&s, "zz", 123)
	//b ,_= setKeyValue(b, "a", "c")
	//b ,_= setKeyValue(b, "next", s)
	//bs,err:=json.Marshal(b)
	//fmt.Println(string(bs),err)
	//
	//return
	r := gin.Default()
	//Session profile
	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions(configure.SessionName, store))

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

		//查询用户信息不用登录
		user.GET("", handler.CheckLoginStatus, handler.QueryUser)
	}
	comment := r.Group("/comment", handler.RequestMiddle, handler.Except)
	{
		comment.POST("", handler.CheckLoginStatus, handler.InsertComment)

		comment.GET(("/:article_id"), handler.QueryComment)
	}
	//解决CORS问题
	r.OPTIONS("/*all", handler.RequestMiddle, handler.Except)
	r.Run(":8080")
}
