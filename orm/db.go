package orm

import (
	"blog_0/configure"
	"blog_0/proerror"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	_db, err := gorm.Open("mysql", configure.GetDSN())
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	} else {
		_db.DB().SetMaxOpenConns(100)
		_db.DB().SetMaxIdleConns(20)
		_db.LogMode(true)
		db = _db
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
	return db
}
