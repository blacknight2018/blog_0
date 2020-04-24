package upfileDao

import (
	"blog_0/orm"
	"time"
)

type UpFile struct {
	Fid                int       `gorm:"column:fid;unique_index;PRIMARY_KEY;"`
	ContentDisposition string    `gorm:"column:content_disposition;"`
	ContentType        string    `gorm:"column:content_type;"`
	FMd5               string    `gorm:"column:fmd5;"`
	CreateTime         time.Time `gorm:"column:create_time;-;"`
}

func (t UpFile) TableName() string {
	return "upfiles"
}

func (t *UpFile) InsertUploadFile() bool {
	err := orm.GetDB().Create(t).Error
	if err == nil {
		return true
	}
	return false
}

func (t *UpFile) QueryGetFile() bool {
	err := orm.GetDB().First(t).Error
	if err == nil {
		return true
	}
	return false
}
