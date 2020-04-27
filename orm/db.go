package orm

import (
	"blog_0/configure"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	_db, err := gorm.Open("mysql", configure.GetDSN())
	if err == nil {
		_db.DB().SetMaxOpenConns(100)
		_db.DB().SetMaxIdleConns(20)
		_db.LogMode(false)
		db = _db
	}
}

func GetDB() *gorm.DB {
	return db
}
