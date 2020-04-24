package utilsDao

import (
	"blog_0/configure"
	"blog_0/orm"
	"github.com/jinzhu/gorm"
)

func AddSelectFiledList(db *gorm.DB, filed []string) *gorm.DB {
	for i := 0; i < len(filed); i++ {
		db = AddSelectFiled(db, filed[i])
	}
	return db
}

func AddSelectFiled(db *gorm.DB, filed string) *gorm.DB {
	//这里既然
	s, err := db.Get(configure.SelectFiledKeyName)
	if s == nil && err == false {
		var s2 []string
		s2 = append(s2, filed)
		return db.Set(configure.SelectFiledKeyName, s2)
	} else if err {
		s2 := s.([]string)
		s2 = append(s2, filed)
		return db.Set(configure.SelectFiledKeyName, s2)
	}

	return db
}
func SetDbSelect(db *gorm.DB) *gorm.DB {
	s, err := db.Get(configure.SelectFiledKeyName)
	if err {
		s2 := s.([]string)
		var s3 string
		for i := 0; i < len(s2); i++ {
			s3 += s2[i]
			if i < len(s2)-1 {
				s3 += ","
			}
		}
		return db.Select(s3)
	}
	return db
}

func SetLimit(db *gorm.DB, offset int, limit int) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Limit(limit).Offset(offset)
}

func OrderByID(db *gorm.DB, primaryId string, order string) *gorm.DB {
	if db == nil {
		db = orm.GetDB()
	}
	return db.Order(primaryId + " " + order)
}
