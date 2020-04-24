package articleDao

import (
	"blog_0/orm"
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

func (t *Article) InsertArticle() bool {
	now := time.Now()
	t.CreateTime = &now
	t.LastTime = &now
	err := orm.GetDB().Create(t).Error
	if err != nil {
		return false
	}
	return true
}
func (t *Article) DeleteArticle() bool {
	err := orm.GetDB().Delete(t).Error
	if err != nil {
		return false
	}
	return true
}

func (t *Article) QueryDetail() bool {
	err := orm.GetDB().First(t).Error
	if err != nil {
		return false
	}
	return true
}

func (t *Article) ChangeSaveArticle() bool {
	err := orm.GetDB().Save(&t).Error
	if err != nil {
		return false
	}
	return true
}

func QueryGetResult(db *gorm.DB) ([]Article, bool) {
	var articles []Article
	err := db.Find(&articles).Error
	if err != nil {
		return nil, false
	}
	return articles, true
}

func QueryPrimaryID() string {
	return "id"
}
