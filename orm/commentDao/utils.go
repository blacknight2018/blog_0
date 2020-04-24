package commentDao

import (
	"blog_0/orm"
	"github.com/jinzhu/gorm"
)

func SetDestArticleId(db *gorm.DB, articleId int) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Where("article_id = ?", articleId)
}
