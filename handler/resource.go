package handler

import (
	"blog_0/configure"
	"blog_0/fileio"
	"blog_0/orm/upfile"
	"blog_0/proerror"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func InsertSingleFileUpload(context *gin.Context) {
	form, err := context.FormFile("file")
	if err != nil || form.Size == 0 {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.FileEmpty})
	}
	file, err := form.Open()
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
	md5ID := fileio.SaveFile(bytes)
	//保持数据一致性，当以下操作发生错误时，删除掉文件,向上抛出异常
	defer func() {
		if err := recover(); err != nil {
			fileio.RemoveFile(md5ID)
			panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
		}
	}()
	ContentDisposition := form.Header.Get("content-disposition")
	ContentType := form.Header.Get("content-type")
	r := upfile.UpFile{
		ContentDisposition: ContentDisposition,
		ContentType:        ContentType,
		FMd5:               md5ID,
	}
	r.InsertFile()
	r.GetFile()
	//过滤字段
	//
	context.Set(configure.ContextFiledName, r)
	return
}

func QueryFile(context *gin.Context) {
	fid := context.Param("id")
	fidInt, err := strconv.Atoi(fid)
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.ParamError})
	}
	r := upfile.UpFile{
		Fid: fidInt,
	}
	r.GetFile()
	MD5Id := r.FMd5
	bytes := fileio.ReadFile(MD5Id)
	context.Header("content-type", r.ContentType)
	context.Header("content-disposition", r.ContentDisposition)
	context.Writer.Write(bytes)
	context.Writer.Flush()

}
