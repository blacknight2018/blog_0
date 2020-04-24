package handler

import (
	"blog_0/configure"
	"blog_0/conversation"
	"blog_0/orm/articleDao"
	"blog_0/orm/utilsDao"
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

func QueryArticle(context *gin.Context) {
	limit := context.DefaultQuery("limit", "10")
	offset := context.DefaultQuery("offset", "0")
	order := context.DefaultQuery("order", "desc")
	flag := context.DefaultQuery("flag", "")
	filed := context.QueryArray("filed")
	limitInt, err := strconv.Atoi(limit)
	offsetInt, err2 := strconv.Atoi(offset)

	//default filed
	defaultFiled := [...]string{"id", "title", "author", "description", "create_time", "last_time"}
	if err == nil && err2 == nil {
		r := utilsDao.OrderByID(nil, articleDao.QueryPrimaryID(), order)
		if flag == "len" {
			r = utilsDao.AddSelectFiled(r, "id")
		} else {
			r = utilsDao.SetLimit(r, offsetInt, limitInt)
			r = utilsDao.AddSelectFiledList(r, defaultFiled[:])
			r = utilsDao.AddSelectFiledList(r, filed)
		}
		r = utilsDao.SetDbSelect(r)
		ret, ok := articleDao.QueryGetResult(r)
		if !ok {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		context.Set(configure.ContextFiledName, ret)
		return
	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
}
func QueryArticleDetail(context *gin.Context) {
	var id = context.Param("id")
	var idInt, err = strconv.Atoi(id)
	if err == nil {
		var article = articleDao.Article{
			Id: idInt,
		}
		if !article.QueryDetail() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		context.Set(configure.ContextFiledName, article)
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
}
func InsertArticle(context *gin.Context) {
	var ret string
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		title := gjson.Get(json, "title").String()
		content := gjson.Get(json, "content").String()
		description := gjson.Get(json, "description").String()
		img := gjson.Get(json, "view_img").String()
		us := conversation.GetSessionUser(context)

		//fmt.Println(img)
		//检查空字段
		if false == checkParamsSafeStringNotEmpty(title, content, description, img) {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
		}

		article := articleDao.Article{
			Title:       title,
			Author:      strconv.Itoa(us.Uid),
			Content:     content,
			Description: description,
			ViewImg:     img,
		}
		if !article.InsertArticle() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		context.Set(configure.ContextFiledName, ret)

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
}

func DeleteArticle(context *gin.Context) {
	var ret string
	var id string = context.Param("id")
	var idInt, err = strconv.Atoi(id)
	if err == nil {
		var article = articleDao.Article{
			Id: idInt,
		}
		if !article.DeleteArticle() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		context.Set(configure.ContextFiledName, ret)
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
}

func ChangeArticle(context *gin.Context) {
	var id string = context.Param("id")
	var idInt, err = strconv.Atoi(id)
	bs, err2 := context.GetRawData()
	if err == nil && err2 == nil {
		json := string(bs)

		content := gjson.Get(json, "content").String()

		var article = articleDao.Article{
			Id: idInt,
		}
		if !article.QueryDetail() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		article.Content = content
		if !article.ChangeSaveArticle() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
}
