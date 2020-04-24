package commentDao

import (
	"blog_0/orm"
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	Cid         int        `json:"cid" gorm:"column:cid;unique_index;PRIMARY_KEY;"`
	AncestorCid int        `json:"ancestor_cid" gorm:"column:ancestor_cid;"`
	Content     string     `json:"content,omitempty" gorm:"column:content;"`
	ReplyToCId  int        `json:"replyto_cid" gorm:"column:replyto_cid;"`
	ArticleId   int        `json:"article_id,omitempty" gorm:"column:article_id;"`
	UserId      int        `json:"uid,omitempty" gorm:"column:uid;"`
	CreateTime  *time.Time `json:"create_time,omitempty" gorm:"column:create_time;-"`
	LastTime    *time.Time `json:"last_time,omitempty" gorm:"column:last_time;-"`
	AuthorName  string     `json:"name" gorm:"-"`
	AuthorHead  string     `json:"avatar" gorm:"-"`
}

func (t Comment) TableName() string {
	return "Comments"
}

func (t *Comment) InsertComment() bool {
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
		r.QueryGetDetail()
		if r.AncestorCid == 0 {
			ancestorCid = replyToCid
		} else {
			ancestorCid = r.AncestorCid
		}
	}
	t.AncestorCid = ancestorCid
	err := orm.GetDB().Create(t).Error
	if err == nil {
		return true
	}
	return false
}

func (t *Comment) QueryGetDetail() bool {
	err := orm.GetDB().First(t).Error
	if err == nil {
		return true
	}
	return false
}

func QueryGetResult(db *gorm.DB) ([]Comment, bool) {
	var comments []Comment
	err := db.Find(&comments).Error
	if err == nil {
		return comments, true
	}
	return nil, false
}

func QueryPrimaryID() string {
	return "cid"
}
