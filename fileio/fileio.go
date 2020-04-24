package fileio

import (
	"blog_0/configure"
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
func SaveFile(FileData []byte) (string, bool) {
	md5ID := md5V()
	FileName := configure.GetLocalFileDir() + md5ID
	err := ioutil.WriteFile(FileName, FileData, 0777)

	if err != nil {
		return "", false
	}
	return md5ID, true
}

func ReadFile(md5ID string) ([]byte, bool) {
	FileName := configure.GetLocalFileDir() + md5ID
	bytes, err := ioutil.ReadFile(FileName)
	if err != nil {
		return nil, false
	}
	return bytes, true
}

func RemoveFile(md5ID string) bool {
	FileName := configure.GetLocalFileDir() + md5ID
	err := os.Remove(FileName)
	if err != nil {
		return false
	}
	return true
}
