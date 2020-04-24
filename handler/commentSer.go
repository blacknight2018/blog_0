package handler

import (
	"blog_0/configure"
	"blog_0/conversation"
	"blog_0/orm/commentDao"
	"blog_0/orm/userDao"
	"blog_0/orm/utilsDao"
	"blog_0/proerror"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"strconv"
)

func InsertComment(context *gin.Context) {
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		content := gjson.Get(json, "content").String()
		replyToId := gjson.Get(json, "replyto_cid").String()
		articleId := gjson.Get(json, "article_id").String()

		if replyToId == "" {
			replyToId = "0"
		}
		replyToIdInt, err2 := strconv.Atoi(replyToId)
		articleIdInt, err3 := strconv.Atoi(articleId)
		us := conversation.GetSessionUser(context)
		if content != "" && err2 == nil && err3 == nil {
			c := commentDao.Comment{
				Content:    content,
				ReplyToCId: replyToIdInt,
				ArticleId:  articleIdInt,
				UserId:     us.Uid,
			}
			c.InsertComment()
			context.Set(configure.ContextFiledName, configure.ContextEmptyFiled)
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
		}
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func QueryComment(context *gin.Context) {
	articleId := context.Param("article_id")
	limit := context.DefaultQuery("limit", "50")
	offset := context.DefaultQuery("offset", "0")
	order := context.DefaultQuery("order", "asc")
	flag := context.DefaultQuery("flag", "")

	//default filed
	defaultFiled := [...]string{"cid", "ancestor_cid", "content", "last_time", "uid", "replyto_cid"}

	//
	limitInt, err := strconv.Atoi(limit)
	offsetInt, err2 := strconv.Atoi(offset)
	articleIdInt, err3 := strconv.Atoi(articleId)
	if err == nil && err2 == nil && err3 == nil {
		r := utilsDao.OrderByID(nil, commentDao.QueryPrimaryID(), order)
		//只返回目标Article
		r = commentDao.SetDestArticleId(r, articleIdInt)
		//只返回cid
		if flag == "len" {
			r = utilsDao.AddSelectFiled(r, "cid")
		} else {
			r = utilsDao.SetLimit(r, offsetInt, limitInt)
			r = utilsDao.AddSelectFiledList(r, defaultFiled[:])
		}
		r = utilsDao.SetDbSelect(r)
		ret := commentDao.QueryGetResult(r)
		//添加上name字段
		for i := 0; i < len(ret); i++ {
			com := ret[i]
			u := userDao.User{
				Uid: com.UserId,
			}
			u.QueryGetUser()

			ret[i].AuthorName = u.User
			ret[i].AuthorHead = u.AvatarUrl
		}

		context.Set(configure.ContextFiledName, ret)
		return

	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})

}
