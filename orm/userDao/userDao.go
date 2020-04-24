package userDao

import (
	"blog_0/orm"
	"blog_0/proerror"
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

func (t *User) InsertUser() {
	now := time.Now()
	t.CreateTime = &now
	err := orm.GetDB().Create(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

/* 可以改的更详细一点:用户不存在 密码错误 */
func (t *User) QueryCheckUser() {
	err := orm.GetDB().Where("user = ? and password = ?", t.User, t.PassWord).Take(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorOpera, ErrorCode: proerror.LoginFiled})
	}
}

/* 根据主键ID 获取信息 */
func (t *User) QueryGetUser() {
	err := orm.GetDB().First(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
