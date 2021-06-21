package xjutils

import (
	"gorm.io/gorm"
)

type QueryWrapper struct {
	Where   string
	Args    []interface{}
	Cols    string
	Orderby string
	Current int
	Size    int
}

func (this QueryWrapper) Set(db *gorm.DB) *gorm.DB {
	if this.Where != "" {
		if len(this.Args) > 0 {
			db = db.Where(this.Where, this.Args...)
		} else {
			db = db.Where(this.Where)
		}
	}
	if this.Cols != "" {
		db = db.Select(this.Cols)
	}
	if this.Orderby != "" {
		db = db.Order(this.Orderby)
	}
	start := (this.Current - 1) * this.Size
	if start > 0 {
		db = db.Offset(start)
	}
	if this.Size > 0 {
		db = db.Limit(this.Size)
	}
	return db
}
