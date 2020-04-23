package handler

import (
	"blog_0/conversation"
	"blog_0/orm/comment"
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
			c := comment.Comment{
				Content:    content,
				ReplyToCId: replyToIdInt,
				ArticleId:  articleIdInt,
				UserId:     us.Uid,
			}
			c.InsertComment()
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
		}
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func QueryComment(context *gin.Context) {

	/*
		articleId :=context.Query("article_id")
		limit := context.DefaultQuery("limit", "10")
		offset := context.DefaultQuery("offset", "0")
		order := context.DefaultQuery("order", "desc")
		flag := context.DefaultQuery("flag", "")
		limitInt, err := strconv.Atoi(limit)
		offsetInt, err2 := strconv.Atoi(offset)


		if err == nil && err2 == nil {

		}
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	*/
}
