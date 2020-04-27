package out

import (
	"blog_0/orm/upfileDao"
	"strings"
)

func getNameFromDisposition(content string) string {
	pos := strings.Index(content, "filename=")
	pos += len("filename=")
	return content[pos+1 : (len(content) - 1)]
}
func QueryFileName(fid int) (string, bool) {
	f := upfileDao.UpFile{
		Fid: fid,
	}
	if f.QueryGetFile() {
		return getNameFromDisposition(f.ContentDisposition), true
	}
	return "", false
}
