package articleSer

import (
	"blog_0/configure"
	"blog_0/handler/resourceSer/out"
	"blog_0/handler/userSer/out/conversation"
	"blog_0/handler/utils"
	"blog_0/orm/articleDao"
	"blog_0/orm/utilsDao"
	"blog_0/proerror"
	"github.com/bennyscetbun/jsongo"
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

// @查询文章
// Name will print hello name
// @Summary 查询文章数量和获取一些文章
// @Description 查询文章,用在后台管理时显示所有数量
// @Accept json
// @Produce  json
// @Param limit query int false "每次获取的数量" default(10)
// @Param offset query int false "每次获取的数量" default(0)
// @Param order query string false "每次获取的数量" default(desc)
// @Param flag query string false "标志位有[len]"
// @Param filed query string false "额外添加的字段"
// @Router /article [get]
// @Success 200 {array} articleDao.Article string "返回一个数组"
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
		r := utilsDao.SetOrderByID(nil, articleDao.QueryPrimaryID(), order)
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
		utils.SetSuccessRetObjectToJSONWithThrowException(context, ret)
		return
	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
}

// @查询文章内容
// Name will print hello name
// @Summary 获取文章的详细内容
// @Description 标题 作者id 内容正文 和附件
// @Accept json
// @Param id path int true "文章ID"
// @Produce  json
// @Router /article/{id}/detail [get]
// @Success 200 {object} articleDao.Article string "此外还包含文件列表数组"
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
		//查看文章内容时，文章封面和描述不需要了
		article.ViewImg = ""
		article.Description = ""

		//查询出文章附带的文件列表返回
		node := jsongo.Node{}
		node.UnmarshalJSON([]byte(utils.GetJsonFromObjectWithThrowException(article)))

		//文件结构
		type fileStruct struct {
			Fid  string `json:"fid"`
			Name string `json:"name"`
		}
		var fs []fileStruct
		for _, v := range gjson.Parse(article.File).Array() {
			fileName, _ := out.QueryFileName(int(v.Int()))
			fs = append(fs, fileStruct{
				v.String(),
				fileName,
			})
		}
		//

		//添加上这个结点
		node.At("file").Val(fs)

		context.Set(configure.ContextFiledName, utils.GetJsonFromNodeObjectParseWithThrowException(&node))

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
}

type swag0 struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	View_img    string `json:"view_img"`
	File        string `json:"file"`
}

// @添加文章
// Name will print hello name
// @Summary 添加文章
// @Description 插入文章到数据库，需要登陆
// @Accept json
// @Produce  json
// @Param title body swag0  true "参数正文"
// @Router /article [POST]
// @Success 200
func InsertArticle(context *gin.Context) {
	var ret string
	bs, err := context.GetRawData()
	if err == nil {
		json := string(bs)
		title := gjson.Get(json, "title").String()
		content := gjson.Get(json, "content").String()
		description := gjson.Get(json, "description").String()
		img := gjson.Get(json, "view_img").String()
		file := gjson.Get(json, "file").String()
		us := conversation.GetSessionUser(context)
		content = utils.Base64String(content)
		title = utils.Base64String(title)
		description = utils.Base64String(description)
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
			File:        file,
		}
		if !article.InsertArticle() {
			panic(proerror.PanicError{
				ErrorType: proerror.ErrorOpera,
				ErrorCode: proerror.UnknownError,
			})
		}
		utils.SetSuccessRetObjectToJSONWithThrowException(context, ret)

	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
}

// @删除文章
// Name will print hello name
// @Summary 删除文章
// @Description 删除文章，需要登陆
// @Param id path int true "文章ID"
// @Router /article/{id} [DELETE]
// @Success 200
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
		utils.SetSuccessRetObjectToJSONWithThrowException(context, ret)
		//context.Set(configure.ContextFiledName, ret)
	} else {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
}

// @修改文章
// Name will print hello name
// @Summary 修改文章
// @Description 修改文章，需要登陆
// @Accept json
// @Param id path int true "文章ID"
// @Param content formData string true "修改后的文章正文"
// @Router /article/{id} [PUT]
// @Success 200
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
