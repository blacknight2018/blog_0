package orm

import (
	"blog_0/proerror"
	"time"
)

type User struct {
	Uid        int        `json:"uid" gorm:"column:uid;unique_index;PRIMARY_KEY"`
	User       string     `json:"user" gorm:"column:user;"`
	PassWord   string     `gorm:"column:password;"`
	Type       int        `gorm:"column:type"`
	AvatarUrl  string     `json:"avatar" gorm:"column:avatar"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
}

func (t User) TableName() string {
	return "users"
}

func (t *User) InsertUser() {
	now := time.Now()
	t.CreateTime = &now
	err := GetDB().Create(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func (t *User) CheckUser() {
	err := GetDB().Where("user = ? and password = ?", t.User, t.PassWord).Take(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.LoginFiled})
	}
}
