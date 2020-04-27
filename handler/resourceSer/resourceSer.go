package resourceSer

import (
	"blog_0/fileio"
	"blog_0/handler/utils"
	"blog_0/orm/upfileDao"
	"blog_0/proerror"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func InsertSingleFileUpload(context *gin.Context) {
	form, err := context.FormFile("file")
	if err != nil || form.Size == 0 {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
	file, err := form.Open()
	defer file.Close()

	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
	md5ID, ok := fileio.SaveFile(bytes)
	if ok {
		ContentDisposition := form.Header.Get("content-disposition")
		ContentType := form.Header.Get("content-type")
		r := upfileDao.UpFile{
			ContentDisposition: ContentDisposition,
			ContentType:        ContentType,
			FMd5:               md5ID,
		}
		ok = r.InsertUploadFile()
		if ok {
			if r.QueryGetFile() {
				utils.SetSuccessRetObjectToJSONWithThrowException(context, r)
				return
			}
		}

	}
	fileio.RemoveFile(md5ID)
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.UnknownError})
}

func QueryFile(context *gin.Context) {
	fid := context.Param("id")
	fidInt, err := strconv.Atoi(fid)
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
	r := upfileDao.UpFile{
		Fid: fidInt,
	}
	if !r.QueryGetFile() {
		panic(proerror.PanicError{
			ErrorType: proerror.ErrorOpera,
			ErrorCode: proerror.UnknownError,
		})
	}
	MD5Id := r.FMd5
	bytes, ok := fileio.ReadFile(MD5Id)

	if ok {
		context.Header("content-type", r.ContentType)
		context.Header("content-disposition", r.ContentDisposition)
		context.Writer.Write(bytes)
		context.Writer.Flush()
		return
	}
	panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})

}
