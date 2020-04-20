package handler

import (
	"blog_0/configure"
	"blog_0/orm"
	"blog_0/proerror"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"strconv"
)

func checkParamsSafeStringNotEmpty(args ...string) bool {
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			return false
		}
	}
	return true
}

//@Accept json
//@Produce json
//@Summary Query
//@Description Query Articles
//@Router /article/
//@Param limit offset order flag
func Query(context *gin.Context) {
	limit := context.DefaultQuery("limit", "10")
	offset := context.DefaultQuery("offset", "0")
	order := context.DefaultQuery("order", "asc")
	flag := context.DefaultQuery("flag", "")
	limitInt, err := strconv.Atoi(limit)
	offsetInt, err2 := strconv.Atoi(offset)
	if err == nil && err2 == nil {

		r := orm.OrderByIDDesc(nil, order)
		if flag == "len" {
			r = orm.SelectOnlyIdField(nil)
		} else {
			r = orm.GetArticleListLimits(r, offsetInt, limitInt)
			r = orm.SelectPreviewField(r)
		}

		ret := orm.GetResult(r)
		if err == nil {
			context.Set(configure.ContextFiledName, ret)
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
		}
		return
	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
}
func QueryDetail(context *gin.Context) {
	var id string = context.Param("id")
	var idInt, err = strconv.Atoi(id)
	if err == nil {
		var article = orm.Article{
			Id: idInt,
		}
		article.GetDetail()
		context.Set(configure.ContextFiledName, article)
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func Insert(context *gin.Context) {
	var ret string
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		author := gjson.Get(json, "author").String()
		title := gjson.Get(json, "title").String()
		content := gjson.Get(json, "content").String()
		description := gjson.Get(json, "description").String()
		img := gjson.Get(json, "view_img").String()

		//fmt.Println(img)
		//检查空字段
		if false == checkParamsSafeStringNotEmpty(author, title, content, description) {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.FieldEmpty})
		}

		article := orm.Article{
			Title:       title,
			Author:      author,
			Content:     content,
			Description: description,
			ViewImg:     img,
		}
		article.InsertArticle()
		context.Set(configure.ContextFiledName, ret)

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func Delete(context *gin.Context) {
	var ret string
	var id string = context.Param("id")
	var idInt, err = strconv.Atoi(id)
	if err == nil {
		var article = orm.Article{
			Id: idInt,
		}
		article.DeleteArticle()
		context.Set(configure.ContextFiledName, ret)
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func Change(context *gin.Context) {
	var id string = context.Param("id")
	var idInt, err = strconv.Atoi(id)
	bs, err2 := context.GetRawData()
	if err == nil && err2 == nil {
		json := string(bs)

		content := gjson.Get(json, "content").String()

		var article = orm.Article{
			Id: idInt,
		}
		article.GetDetail()
		article.Content = content
		article.Save()

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
