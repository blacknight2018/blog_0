package commentSer

import (
	"blog_0/configure"
	"blog_0/handler/userSer/out"
	"blog_0/handler/userSer/out/conversation"
	"blog_0/handler/utils"
	"blog_0/orm/commentDao"
	"blog_0/orm/utilsDao"
	"blog_0/proerror"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"strconv"
)

type swag1 struct {
	Content     string `json:"content"`
	Replyto_cid string `json:"replyto_cid"`
	Article_id  string `json:"article_id"`
}

// @添加文章
// Name will print hello name
// @Summary 发表回复
// @Description 回复
// @Accept json
// @Produce  json
// @Param title body swag1  true "参数正文"
// @Router /comment [POST]
// @Success 200
func InsertComment(context *gin.Context) {
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		content := gjson.Get(json, "content").String()
		replyToId := gjson.Get(json, "replyto_cid").String()
		articleId := gjson.Get(json, "article_id").String()
		content = utils.Base64String(content)
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
			utils.SetSuccessRetObjectToJSONWithThrowException(context, configure.ContextEmptyFiled)
		} else {
			panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
		}
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.UnknownError})
	}
}

type swag2 struct {
	commentDao.Comment
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

// @查询回复
// Name will print hello name
// @Summary 获取回复
// @Description 获取回复内容，头像，正文，名称
// @Accept json
// @Produce  json
// @Param article_id path int   true "文章id"
// @Param limit query int   false "获取数量"
// @Param offset query int   false "开始偏移"
// @Param order query string   false "顺序"
// @Param flag query string   false "标志位 [len]"
// @Router /comment/{article_id} [GET]
// @Success 200 {object} swag2
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
		r := utilsDao.SetOrderByID(nil, commentDao.QueryPrimaryID(), order)
		//只返回目标Article
		r = commentDao.SetDestArticleId(r, articleIdInt)
		//只返回cid
		if flag == "len" {
			r = utilsDao.AddSelectFiled(r, "cid")
			r = utilsDao.SetDbSelect(r)
			ret, ok := commentDao.QueryGetResult(r)
			if ok {
				utils.SetSuccessRetObjectToJSONWithThrowException(context, ret)
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
		retJson := utils.GetJsonFromObjectWithThrowException(ret)
		jsonObj := utils.GetNodeObjectFromJsonWithThrowException(retJson)
		for i := 0; i < len(ret); i++ {
			com := ret[i]
			userName, err := out.GetUserName(com.UserId)
			userAvatar, err2 := out.GetUserAvatar(com.UserId)
			if !err || !err2 {
				panic(proerror.PanicError{
					ErrorType: proerror.ErrorOpera,
					ErrorCode: proerror.UnknownError,
				})
			}
			jsonObj.At(i).At("name").Val(userName)
			jsonObj.At(i).At("avatar").Val(userAvatar)
		}
		context.Set(configure.ContextFiledName, utils.GetJsonFromNodeObjectParseWithThrowException(&jsonObj))
		return

	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})

}
