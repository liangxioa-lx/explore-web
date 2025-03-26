package common

import (
	"gorm.io/gorm"
)

type Page struct {
	PageNum  int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

func Paginate(req *Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req == nil || (req.PageSize == 0 && req.PageNum == 0) {
			return db
		}
		page := req.PageNum
		if page == 0 {
			page = 1
		}
		pageSize := req.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func CheckPageParams(page Page) bool {
	if page.PageNum == 0 || page.PageSize == 0 {
		return false
	}
	return true
}
