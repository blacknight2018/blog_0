package handler

import (
	"blog_0/configure"
	"blog_0/conversation"
	"blog_0/orm/comment"
	"blog_0/orm/user"
	"blog_0/proerror"
	"fmt"
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
			c := comment.Comment{
				Content:    content,
				ReplyToCId: replyToIdInt,
				ArticleId:  articleIdInt,
				UserId:     us.Uid,
			}
			c.InsertComment()
			context.Set(configure.ContextFiledName, "")
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
		}
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func QueryComment(context *gin.Context) {
	articleId := context.Param("article_id")
	limit := context.DefaultQuery("limit", "10")
	offset := context.DefaultQuery("offset", "0")
	order := context.DefaultQuery("order", "asc")
	flag := context.DefaultQuery("flag", "")

	//
	limitInt, err := strconv.Atoi(limit)
	offsetInt, err2 := strconv.Atoi(offset)
	articleIdInt, err3 := strconv.Atoi(articleId)
	if err == nil && err2 == nil && err3 == nil {
		r := comment.OrderByIDDesc(nil, order)

		//只返回目标Article
		r = comment.SetDestArticleId(r, articleIdInt)
		//只返回id
		if flag == "len" {
			r = comment.SelectOnlyIdField(r)
		} else {
			//根据参数过滤
			r = comment.SetCommentListLimit(r, offsetInt, limitInt)

			//返回部分字段
			r = comment.SelectPreviewField(r)
		}
		ret := comment.GetResult(r)
		//添加上name字段
		for i := 0; i < len(ret); i++ {
			com := ret[i]
			fmt.Println(com.UserId)
			u := user.User{
				Uid: com.UserId,
			}
			u.GetUser()

			ret[i].AuthorName = u.User
			ret[i].AuthorHead = u.AvatarUrl
		}

		context.Set(configure.ContextFiledName, ret)
		return

	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})

}
