package userDao

import (
	"blog_0/orm"
	"time"
)

type User struct {
	Uid        int        `json:"uid" gorm:"column:uid;unique_index;PRIMARY_KEY"`
	User       string     `json:"user" gorm:"column:user;"`
	PassWord   string     `json:"-" gorm:"column:password;"`
	Type       int        `json:"type" gorm:"column:type"`
	AvatarUrl  string     `json:"avatar" gorm:"column:avatar"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
}

func (t User) TableName() string {
	return "users"
}

func (t *User) InsertUser() bool {
	now := time.Now()
	t.CreateTime = &now
	err := orm.GetDB().Create(t).Error
	if err == nil {
		return true
	}
	return false
}

func (t *User) QueryCheckUserPassWord() bool {
	err := orm.GetDB().Where("user = ? and password = ?", t.User, t.PassWord).Take(t).Error
	if err == nil {
		return true
	}
	return false
}

func (t *User) QueryGetUser() bool {
	err := orm.GetDB().First(t).Error
	if err == nil && t.Uid > 0 {
		return true
	}
	return false
}
