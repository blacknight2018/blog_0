package main

import (
	"blog_0/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	//type student struct {
	//	Name string
	//}
	//s1:=student{Name:"a"}
	//s2:=student{Name:"b"}
	//var big []student
	//big = append(big, s1)
	//big = append(big, s2)
	//ss:=utils.JsonParseWithThrowException(big)
	//
	//var array []interface{}
	//json.Unmarshal([]byte(ss),&array)
	//s12:=utils.JsonParseWithThrowException(array[0])
	//s22:=utils.JsonParseWithThrowException(array[1])
	//fmt.Println(s12,s22)
	//fmt.Println(ss)
	//o := jsongo.EmptyJsonObject()
	//o.SetKey("Acc", 22)
	////
	//o2 := jsongo.EmptyJsonObject()
	//
	//////o.SetKey("cc",s1)
	//o2.SetKey("person2", o)
	//array := jsongo.EmptyJsonArrayObject()
	//array.Add(*o2)
	//array.Add(*o2)
	////array.Add(o2)
	//fmt.Println(array.GetJSONString())
	////fmt.Println(array.Get(0))
	////fmt.Println(array.Get(1))
	//array.Get(0).SetKey("aa", 33)
	//array.Get(1).SetKey("aabbb", 33)
	//fmt.Println(array.GetJSONString())

	//j:=jsongo.EmptyJsonObject().SetKey("name","chen")
	//j.SetKey("na22me","chen")
	//j2:=jsongo.EmptyJsonObject().SetKey("name","chen")
	//j.SetKey("1",j2)
	//fmt.Println(j.GetJSONString())
	//return
	//j:=jsongo.EmptyJsonArrayObject().Add(*jsongo.EmptyJsonObject().SetKey("a",1))
	//j.Add(*jsongo.EmptyJsonObject().SetKey("name","chenhuaze"))
	//fmt.Println(j.GetJSONString())
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
	r.Run(":8080")
}
