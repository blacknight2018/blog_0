package comment

import (
	"blog_0/orm"
	"blog_0/proerror"
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	Cid         int        `json:"cid" gorm:"column:cid;unique_index;PRIMARY_KEY;"`
	AncestorCid int        `json:"ancestor_cid,omitempty" gorm:"column:ancestor_cid;"`
	Content     string     `json:"content,omitempty" gorm:"column:content;"`
	ReplyToCId  int        `json:"replyto_id,omitempty" gorm:"column:replyto_cid;"`
	ArticleId   int        `json:"article_id,omitempty" gorm:"column:article_id;"`
	UserId      int        `json:"uid,omitempty" gorm:"column:uid;"`
	CreateTime  *time.Time `json:"create_time,omitempty" gorm:"column:create_time;-"`
	LastTime    *time.Time `json:"last_time,omitempty" gorm:"column:last_time;-"`
}

func (t Comment) TableName() string {
	return "Comments"
}

func (t *Comment) InsertComment() {
	now := time.Now()
	t.CreateTime = &now
	t.LastTime = &now

	//使回复构成一个树形结构，并且每个结点都有一个根节点ID
	var ancestorCid int = 0
	//看回复的是否是顶层回复
	replyToCid := t.ReplyToCId
	if replyToCid != 0 {
		r := Comment{
			Cid: replyToCid,
		}
		r.GetDetail()
		if r.AncestorCid == 0 {
			ancestorCid = replyToCid
		} else {
			ancestorCid = r.AncestorCid
		}
	}
	t.AncestorCid = ancestorCid
	err := orm.GetDB().Create(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func (t *Comment) GetDetail() {
	err := orm.GetDB().First(t).Error
	if err != nil {
		panic(proerror.PanicError{ErrorType: proerror.ErrorIo})
	}
}

func OrderByIDDesc(db *gorm.DB, order string) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Order("cid " + order)
}

func SelectOnlyIdField(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Select("cid")
}

func SetCommentListLimits(db *gorm.DB, offset int, limit int) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Limit(limit).Offset(offset)
}

func SelectPreviewField(db *gorm.DB) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Select("cid,ancestor_cid,content,last_time")
}

func SetDestArticleId(db *gorm.DB, articleId int) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Where("article_id = ?", articleId)
}

func GetResult(db *gorm.DB) []Comment {
	var comments []Comment
	if db != nil {
		db.Find(&comments)
	}
	return comments
}
