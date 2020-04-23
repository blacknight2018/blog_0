package article

import (
	"blog_0/orm"
	"blog_0/proerror"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Id          int        `json:"id" gorm:"column:id;unique_index;PRIMARY_KEY"`
	Title       string     `json:"title,omitempty" gorm:"column:title"`
	Author      string     `json:"author,omitempty" gorm:"column:author"`
	Content     string     `json:"content,omitempty" gorm:"column:content"`
	Description string     `json:"description,omitempty" gorm:"column:description"`
	Like        int        `json:"like,omitempty" gorm:"column:like;DEFAULT:0"`
	CreateTime  *time.Time `json:"create_time,omitempty" gorm:"column:create_time;"`
	LastTime    *time.Time `json:"last_time,omitempty" gorm:"column:last_time;-"`
	ViewImg     string     `json:"view_img,omitempty" gorm:"column:view_img;"`
}

func (t Article) TableName() string {
	return "articles"
}

func (t *Article) InsertArticle() {
	now := time.Now()
	t.CreateTime = &now
	t.LastTime = &now
	err := orm.GetDB().Create(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func (t *Article) DeleteArticle() {
	err := orm.GetDB().Delete(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func (t *Article) GetDetail() {
	err := orm.GetDB().First(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func (t *Article) SaveArticle() {
	err := orm.GetDB().Save(&t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}
func SetArticleListLimits(db *gorm.DB, offset int, limit int) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Limit(limit).Offset(offset)
}

func OrderByIDDesc(db *gorm.DB, order string) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Order("id " + order)
}
func SelectPreviewField(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Select("id,title,author,description,create_time,last_time,view_img")
}

func SelectOnlyIdField(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Select("id")
}

func GetResult(db *gorm.DB) []Article {
	var articles []Article
	if db != nil {
		db.Find(&articles)
	}
	return articles
}
