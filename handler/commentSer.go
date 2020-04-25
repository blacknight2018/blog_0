package handler

import (
	"blog_0/configure"
	"blog_0/conversation"
	"blog_0/handler/utils"
	"blog_0/orm/commentDao"
	"blog_0/orm/userDao"
	"blog_0/orm/utilsDao"
	"blog_0/proerror"
	"encoding/json"
	"github.com/donnie4w/json4g"
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
			if !c.InsertComment() {
				panic(proerror.PanicError{
					ErrorType: proerror.ErrorOpera,
					ErrorCode: proerror.UnknownError,
				})
			}
			//context.Set(configure.ContextFiledName, configure.ContextEmptyFiled)
			utils.SetRetObjectToJSONWithThrowException(context, configure.ContextEmptyFiled)
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
		}
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.UnknownError})
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
			r = utilsDao.SetDbSelect(r)
			ret, ok := commentDao.QueryGetResult(r)
			if ok {
				utils.SetRetObjectToJSONWithThrowException(context, ret)
				return
			}
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
			return
		} else {
			r = utilsDao.SetLimit(r, offsetInt, limitInt)
			r = utilsDao.AddSelectFiledList(r, defaultFiled[:])
		}
		r = utilsDao.SetDbSelect(r)
		ret, ok := commentDao.QueryGetResult(r)
		if !ok {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		//添加上name字段
		bytes, err := json.Marshal(&ret)
		if err != nil {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		node, err := json4g.LoadByString(string(bytes))
		for i := 0; i < len(ret); i++ {
			com := ret[i]
			u := userDao.User{
				Uid: com.UserId,
			}
			if !u.QueryGetUser() {
				panic(proerror.PanicError{
					ErrorType: proerror.ErrorOpera,
					ErrorCode: proerror.UnknownError,
				})
			}
			node.ArraysStruct[i].ToJsonNode().AddNode(json4g.NowJsonNode("name", u.User))
			node.ArraysStruct[i].ToJsonNode().AddNode(json4g.NowJsonNode("avatar", u.AvatarUrl))
		}
		context.Set(configure.ContextFiledName, node.ToString())
		return

	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})

}
