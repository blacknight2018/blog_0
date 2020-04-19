package fileio

import (
	"blog_0/configure"
	"blog_0/proerror"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"time"
)

func md5V() string {
	h := md5.New()
	str := time.Now().String()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func SaveFile(FileData []byte) string {
	md5ID := md5V()
	FileName := configure.GetLocalFileDir() + md5ID
	err := ioutil.WriteFile(FileName, FileData, 0777)

	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
	return md5ID
}

func ReadFile(md5ID string) []byte {
	FileName := configure.GetLocalFileDir() + md5ID
	bytes, err := ioutil.ReadFile(FileName)
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
	return bytes
}

func RemoveFile(md5ID string) {
	FileName := configure.GetLocalFileDir() + md5ID
	err := os.Remove(FileName)
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
