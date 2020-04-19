package orm

import (
	"blog_0/proerror"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Id          int       `gorm:"column:id;unique_index;PRIMARY_KEY"`
	Title       string    `gorm:"column:title"`
	Author      string    `gorm:"column:author"`
	Content     string    `gorm:"column:content"`
	Description string    `gorm:"column:description"`
	Like        int       `gorm:"column:like;DEFAULT:0"`
	CreateTime  time.Time `gorm:"column:create_time;"`
	LastTime    time.Time `gorm:"column:last_time;-"`
}

func (t Article) TableName() string {
	return "articles"
}

func (t *Article) InsertArticle() {
	t.CreateTime = time.Now()
	err := GetDB().Create(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func (t *Article) DeleteArticle() {
	err := GetDB().Delete(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func (t *Article) GetDetail() {
	err := GetDB().First(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func (t *Article) Save() {
	err := GetDB().Save(&t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func GetArticleListLimits(db *gorm.DB, offset int, limit int) *gorm.DB {
	if db == nil {
		db = GetDB()
	}
	return db.Limit(limit).Offset(offset)
}

func OrderByIDDesc(db *gorm.DB, order string) *gorm.DB {
	if db == nil {
		db = GetDB()
	}
	return db.Order("id " + order)
}
func SelectPreviewField(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = GetDB()
	}
	return db.Select("id,title,author,description,create_time,last_time")
}
func GetResult(db *gorm.DB) []Article {
	var articles []Article
	if db != nil {
		db.Find(&articles)
	}
	return articles
}
